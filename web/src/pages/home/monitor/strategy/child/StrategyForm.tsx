import { FC, ReactNode, useEffect, useState } from 'react'
import {
    Button,
    Col,
    Form,
    FormInstance,
    Input,
    Row,
    Select,
    Space
} from 'antd'
import PromQLInput, {
    PromValidate,
    formatExpressionFunc
} from '@/components/Prom/PromQLInput.tsx'
import { strategyEditOptions, sverityOptions } from '../options'
import { DeleteOutlined } from '@ant-design/icons'
import AddLabelModal from './AddLabelModal'
import DataForm from '@/components/Data/DataForm/DataForm'
import { DefaultOptionType } from 'antd/es/select'
import { Rule } from 'antd/es/form'
import { StrategyItemType } from '@/apis/home/monitor/strategy/types'
import { Duration } from '@/apis/types'

export type FormValuesType = {
    alert?: string
    annotations?: {
        title: string
        description: string
        [key: string]: string
    }
    duration?: Duration
    dataSource?: DefaultOptionType
    groupId?: number
    lables?: { sverity?: string; [key: string]: string | undefined }
    expr?: string
    restrain?: number[]
    alarmPageIds?: number[]
    categoryIds?: number[]
    remark?: string
    alarmLevelId?: number
    dataSourceId: number
    // 最大抑制时常
    maxSuppress?: Duration
    // 告警通知间隔
    sendInterval?: Duration
    // 是否发送告警通知
    sendRecover?: boolean
}
export interface StrategyFormProps {
    form: FormInstance
    disabled?: boolean
    initialValue?: StrategyItemType
}

export type labelsType = {
    label: string | ReactNode
    name: string
}

let timeout: NodeJS.Timeout
export const StrategyForm: FC<StrategyFormProps> = (props) => {
    const { disabled, form, initialValue } = props

    const [labelFormItemList, setLabelFormItemList] = useState<labelsType[]>([])
    const [annotationFormItemList, setAnnotationFormItemList] = useState<
        labelsType[]
    >([])
    const [addLabelModalOpen, setAddLabelModalOpen] = useState<boolean>(false)
    const [isLabelModalOpen, setIsLabelModalOpen] = useState<boolean>(false)
    const [validatePromQL, setValidatePromQL] = useState<PromValidate>({})

    const dataSource = Form.useWatch<DefaultOptionType>('dataSource', form)

    const buildInitvalue = () => {
        const value = initialValue
        if (!value) {
            form?.resetFields()
            return
        }
        const init: FormValuesType = {
            ...value,
            lables: {
                ...value?.labels,
                sverity: value?.alarmLevelId
                    ? value.alarmLevelId + ''
                    : undefined
            },
            annotations: {
                ...value?.annotations,
                title: value?.annotations?.['title'],
                description: value?.annotations?.['description']
            },
            dataSource: {
                value: value.dataSource?.value,
                label: value.dataSource?.label,
                title: value.dataSource?.endpoint
            },
            restrain: [],
            alert: value?.alert,
            duration: value?.duration,
            alarmLevelId: value?.alarmLevelId,
            alarmPageIds: value?.alarmPageIds,
            expr: value?.expr,
            groupId: value?.groupId,
            categoryIds: value?.categoryIds,
            sendRecover: false
        }
        form?.setFieldsValue(init)
    }

    const fetchValidateExpr = (value?: string) => {
        if (timeout) {
            clearTimeout(timeout)
        }
        if (!dataSource || !dataSource.title) {
            setValidatePromQL({
                help: '请先选择数据源',
                validateStatus: 'error'
            })
            return Promise.resolve()
        }

        if (!value) {
            setValidatePromQL({
                help: '请填写PromQL',
                validateStatus: 'error'
            })
            return Promise.resolve()
        }

        timeout = setTimeout(() => {
            formatExpressionFunc(dataSource?.title, value)
                .then((resp) => {
                    if (resp.status === 'error') {
                        const msg = `[${resp.errorType}] ${resp.error}`
                        setValidatePromQL({
                            help: `[${resp.errorType}] ${resp.error}`,
                            validateStatus: 'error'
                        })
                        return Promise.reject(new Error(msg))
                    }
                    if (resp.status === 'success') {
                        setValidatePromQL({
                            help: '语法校验通过✅',
                            validateStatus: 'success'
                        })
                    }
                    return resp
                })
                .catch((err: any) => {
                    setValidatePromQL({
                        help: `${err}`,
                        validateStatus: 'error'
                    })
                    return err
                })
        }, 200)

        return Promise.resolve()
    }

    const PromQLRule: Rule[] = [
        {
            required: true,
            message: 'PromQL不能为空, 请填写PromQL'
        },
        {
            validator: (_: Rule, value: string) => {
                return fetchValidateExpr(value)
            }
        }
    ]

    const handleAddLabel = (data: labelsType) => {
        if (data.label && data.name) {
            if (isLabelModalOpen) {
                setLabelFormItemList([...labelFormItemList, data])
            } else {
                setAnnotationFormItemList([...annotationFormItemList, data])
            }
            setAddLabelModalOpen(false)
        }
    }

    const handleCloseAddLabelModal = () => {
        setAddLabelModalOpen(false)
    }

    const openAddLabelModal = () => {
        setAddLabelModalOpen(true)
        setIsLabelModalOpen(true)
    }

    const openAddAnnotationModal = () => {
        setAddLabelModalOpen(true)
        setIsLabelModalOpen(false)
    }

    const handleDeleteLabelFormItemListByIndex = (index: number) => {
        setLabelFormItemList(labelFormItemList.filter((_, i) => i !== index))
    }

    const handleDeleteAnnotationFormItemListByIndex = (index: number) => {
        setAnnotationFormItemList(
            annotationFormItemList.filter((_, i) => i !== index)
        )
    }

    useEffect(() => {
        buildInitvalue()
    }, [initialValue])

    return (
        <>
            <AddLabelModal
                open={addLabelModalOpen}
                onCancel={handleCloseAddLabelModal}
                onOk={handleAddLabel}
            />
            <DataForm
                form={form}
                items={strategyEditOptions}
                formProps={{
                    layout: 'vertical',
                    disabled: disabled
                }}
            >
                <Form.Item
                    name="expr"
                    label="PromQL"
                    {...validatePromQL}
                    required
                    tooltip={
                        <div>
                            正确的PromQL表达式,
                            用于完成Prometheus报警规则数据匹配
                        </div>
                    }
                    rules={PromQLRule}
                    dependencies={['dataSource']}
                >
                    <PromQLInput
                        disabled={disabled}
                        pathPrefix={dataSource?.title}
                        formatExpression={true}
                    />
                </Form.Item>
                <Form.Item
                    tooltip={
                        <div>
                            标签: 标签是Prometheus报警规则的附加信息, 例如:
                            告警等级, 告警实例等, 也可以添加自定义标签
                        </div>
                    }
                    label={
                        <span style={{ color: '#E0E2E6' }}>
                            <Button
                                type="primary"
                                size="small"
                                onClick={openAddLabelModal}
                            >
                                告警标签
                            </Button>
                            <span>(可选)</span>
                        </span>
                    }
                >
                    <Row gutter={16}>
                        <Col span={6}>
                            <Form.Item
                                name={['lables', 'sverity']}
                                label="等级(sverity)"
                                rules={[
                                    {
                                        required: true,
                                        message: '请输入告警等级'
                                    }
                                ]}
                            >
                                <Select
                                    placeholder="请选择告警等级"
                                    options={sverityOptions}
                                />
                            </Form.Item>
                        </Col>
                        {labelFormItemList.map((item, index) => {
                            return (
                                <Col span={6} key={index}>
                                    <Form.Item
                                        name={['lables', item.name]}
                                        label={`${item.label}(${item.name})`}
                                        rules={[
                                            {
                                                required: true,
                                                message: `请输入${item.label}`
                                            }
                                        ]}
                                    >
                                        <Space.Compact>
                                            <Input
                                                placeholder={`请输入${item.label}`}
                                            />
                                            <Button
                                                type="primary"
                                                danger
                                                icon={<DeleteOutlined />}
                                                onClick={() =>
                                                    handleDeleteLabelFormItemListByIndex(
                                                        index
                                                    )
                                                }
                                            />
                                        </Space.Compact>
                                    </Form.Item>
                                </Col>
                            )
                        })}
                    </Row>
                </Form.Item>

                <Form.Item
                    tooltip={
                        <div>
                            告警注释: 告警注释是Prometheus报警规则的附加信息,
                            例如: 告警标题, 告警描述等, 也可以添加自定义注释
                        </div>
                    }
                    label={
                        <span style={{ color: '#E0E2E6' }}>
                            <Button
                                type="primary"
                                size="small"
                                onClick={openAddAnnotationModal}
                            >
                                告警注释
                            </Button>
                            <span style={{ color: '#E0E2E6' }}>(可选)</span>
                        </span>
                    }
                >
                    <Form.Item
                        name={['annotations', 'title']}
                        label="告警标题模板"
                        rules={[
                            {
                                required: true,
                                message: '请输入告警标题模板'
                            }
                        ]}
                    >
                        <Input.TextArea placeholder="请输入告警标题模板" />
                    </Form.Item>
                    <Form.Item
                        name={['annotations', 'description']}
                        label="告警内容模板"
                        rules={[
                            {
                                required: true,
                                message: '请输入告警内容模板'
                            }
                        ]}
                    >
                        <Input.TextArea placeholder="请输入告警内容模板" />
                    </Form.Item>
                    {annotationFormItemList.map((item, index) => {
                        return (
                            <Form.Item
                                name={['annotations', item.name]}
                                label={`${item.label}(${item.name})`}
                                rules={[
                                    {
                                        required: true,
                                        message: `请输入${item.label}`
                                    }
                                ]}
                                key={index}
                            >
                                <Space.Compact style={{ width: '100%' }}>
                                    <Input.TextArea
                                        placeholder={`请输入${item.label}`}
                                    />
                                    <Button
                                        type="primary"
                                        danger
                                        icon={<DeleteOutlined />}
                                        onClick={() =>
                                            handleDeleteAnnotationFormItemListByIndex(
                                                index
                                            )
                                        }
                                    />
                                </Space.Compact>
                            </Form.Item>
                        )
                    })}
                </Form.Item>
            </DataForm>
        </>
    )
}
