import { POST } from '@/apis/request'
import type {
    CreateChatGroupRequest,
    CreateChatGroupResponse,
    DeleteChatGroupRequest,
    DeleteChatGroupResponse,
    GetChatGroupDetailRequest,
    GetChatGroupDetailResponse,
    ListChatGroupRequest,
    ListChatGroupResponse,
    SelectChatGroupRequest,
    SelectChatGroupResponse,
    UpdateChatGroupRequest,
    UpdateChatGroupResponse
} from './types'

enum URL {
    LIST = '/api/v1/chat/group/list',
    SELECT = '/api/v1/chat/group/select',
    UPDATE = '/api/v1/chat/group/update',
    DELETE = '/api/v1/chat/group/delete',
    DETAIL = '/api/v1/chat/group/get',
    CREATE = '/api/v1/chat/group/create'
}

const getChatGroupList = (params: ListChatGroupRequest) => {
    return POST<ListChatGroupResponse>(URL.LIST, params)
}

const getChatGroupDetail = (params: GetChatGroupDetailRequest) => {
    return POST<GetChatGroupDetailResponse>(URL.DETAIL, params)
}

const createChatGroup = (params: CreateChatGroupRequest) => {
    return POST<CreateChatGroupResponse>(URL.CREATE, params)
}

const updateChatGroup = (params: UpdateChatGroupRequest) => {
    return POST<UpdateChatGroupResponse>(URL.UPDATE, params)
}

const deleteChatGroup = (params: DeleteChatGroupRequest) => {
    return POST<DeleteChatGroupResponse>(URL.DELETE, params)
}

const getChatGroupSelect = (params: SelectChatGroupRequest) => {
    return POST<SelectChatGroupResponse>(URL.SELECT, params)
}

const chatGroupApi = {
    getChatGroupList,
    getChatGroupDetail,
    createChatGroup,
    updateChatGroup,
    deleteChatGroup,
    getChatGroupSelect
}

export default chatGroupApi
