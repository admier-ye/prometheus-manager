import React, { useEffect, useRef, useState } from 'react'
import RouteBreadcrumb from '@/components/PromLayout/RouteBreadcrumb'
import { DataOption, DataTable, SearchForm } from '@/components/Data'

import {
    ListEndpointRequest,
    PrometheusServerItem,
    defaultListEndpointRequest
} from '@/apis/home/monitor/endpoint/types.ts'
import { ActionKey } from '@/apis/data.ts'
import { Form } from 'antd'
import { HeightLine, PaddingLine } from '@/components/HeightLine'
import endpointApi from '@/apis/home/monitor/endpoint/index.ts'
import {
    columns,
    defaultPadding,
    leftOptions,
    operationItems,
    rightOptions,
    searchItems
} from './options'
import EditEndpointModal from './child/EditEnpointModal'
import { Status } from '@/apis/types'

let timer: NodeJS.Timeout

const Endpoint: React.FC = () => {
    const operationRef = useRef<HTMLDivElement>(null)
    const [queryForm] = Form.useForm()

    const [dataSource, setDataSource] = useState<PrometheusServerItem[]>([])
    const [refresh, setRefresh] = useState<boolean>(false)
    const [loading, setLoading] = useState<boolean>(false)
    const [search, setSearch] = useState<ListEndpointRequest>(
        defaultListEndpointRequest
    )
    const [total, setTotal] = useState<number>(0)
    const [openEditModal, setEditModal] = useState<boolean>(false)
    const [opEndpointId, setOpEndpointId] = useState<number>()

    // 刷新
    const handlerRefresh = () => {
        setRefresh((prev) => !prev)
    }

    const hanleOpenEditModal = (id?: number) => {
        setOpEndpointId(id)
        setEditModal(true)
    }

    const handleOnCloseEditModal = () => {
        setEditModal(false)
        setOpEndpointId(undefined)
    }

    const handleEditModalOnOk = () => {
        handleOnCloseEditModal()
        handlerRefresh()
    }

    const handlerGetData = () => {
        if (timer) {
            clearTimeout(timer)
        }
        setLoading(true)
        timer = setTimeout(() => {
            endpointApi
                .listEndpoint(search)
                .then((data) => {
                    setDataSource(data?.list || [])
                    setTotal(data.page.total)
                })
                .finally(() => {
                    setLoading(false)
                })
        }, 500)
    }

    //操作栏按钮
    const handleOptionClick = (val: ActionKey) => {
        switch (val) {
            case ActionKey.ADD:
                hanleOpenEditModal()
                break
            case ActionKey.REFRESH:
                handlerRefresh()
                break
            case ActionKey.RESET:
                setSearch({
                    keyword: '',
                    page: {
                        curr: 1,
                        size: 10
                    }
                })
                break
        }
    }

    const handlerTablePageChange = (page: number, size: number) => {
        setSearch({
            ...search,
            page: {
                curr: page,
                size: size
            }
        })
    }

    // 处理搜索表单的值变化
    const handlerSearFormValuesChange = (changedValues: any) => {
        setSearch({
            ...search,
            ...changedValues
        })
    }

    const handleChangeStatus = (ids: number[], status: Status) => {
        endpointApi.batchChangeStatus(ids, status).then(() => {
            handlerRefresh()
        })
    }

    const handlerTableAction = (key: ActionKey, item: PrometheusServerItem) => {
        switch (key) {
            case ActionKey.EDIT:
                hanleOpenEditModal(item.id)
                break
            case ActionKey.DELETE:
                break
            case ActionKey.ENABLE:
                handleChangeStatus([item.id], Status.STATUS_ENABLED)
                break
            case ActionKey.DISABLE:
                handleChangeStatus([item.id], Status.STATUS_DISABLED)
                break
            default:
                break
        }
    }

    useEffect(() => {
        handlerRefresh()
    }, [search])

    useEffect(() => {
        handlerGetData()
    }, [refresh])

    return (
        <div className="bodyContent">
            <EditEndpointModal
                endpointId={opEndpointId}
                open={openEditModal}
                onClose={handleOnCloseEditModal}
                onOk={handleEditModalOnOk}
            />
            <div ref={operationRef}>
                <RouteBreadcrumb />
                <HeightLine />
                <SearchForm
                    form={queryForm}
                    items={searchItems}
                    formProps={{
                        onValuesChange: handlerSearFormValuesChange
                    }}
                />
                <HeightLine />
                <DataOption
                    queryForm={queryForm}
                    rightOptions={rightOptions(loading)}
                    leftOptions={leftOptions}
                    action={handleOptionClick}
                />
                <PaddingLine
                    padding={defaultPadding}
                    height={1}
                    borderRadius={4}
                />
            </div>
            <DataTable
                dataSource={dataSource}
                columns={columns}
                operationRef={operationRef}
                total={total}
                loading={loading}
                operationItems={operationItems}
                pageOnChange={handlerTablePageChange}
                action={handlerTableAction}
            />
        </div>
    )
}

export default Endpoint
