package bo

import (
	"prometheus-manager/api"
	"prometheus-manager/app/prom_server/internal/biz/do"
	"prometheus-manager/app/prom_server/internal/biz/vo"
)

type (
	ChatGroupBO struct {
		Id        uint32       `json:"id"`
		Name      string       `json:"name"`
		Status    vo.Status    `json:"status"`
		Remark    string       `json:"remark"`
		CreatedAt int64        `json:"createdAt"`
		UpdatedAt int64        `json:"updatedAt"`
		DeletedAt int64        `json:"deletedAt"`
		Hook      string       `json:"hook"`
		NotifyApp vo.NotifyApp `json:"notifyApp"`
		HookName  string       `json:"hookName"`
		// 消息模板
		Template string `json:"template"`
		// 消息标题
		Title string `json:"title"`
		// 通信密钥
		Secret string `json:"secret"`
	}
)

// ToApi ...
func (b *ChatGroupBO) ToApi() *api.ChatGroup {
	if b == nil {
		return nil
	}
	return &api.ChatGroup{
		Id:        b.Id,
		Name:      b.Name,
		Remark:    b.Remark,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		Hook:      b.Hook,
		Status:    b.Status.Value(),
		App:       b.NotifyApp.Value(),
		HookName:  b.HookName,
		Template:  b.Template,
		Title:     b.Title,
		Secret:    b.Secret,
	}
}

// ToSelectApi ...
func (b *ChatGroupBO) ToSelectApi() *api.ChatGroupSelectV1 {
	if b == nil {
		return nil
	}
	return &api.ChatGroupSelectV1{
		Value:  b.Id,
		App:    b.NotifyApp.Value(),
		Label:  b.HookName,
		Status: b.Status.Value(),
	}
}

func (b *ChatGroupBO) ToModel() *do.PromAlarmChatGroup {
	if b == nil {
		return nil
	}
	return &do.PromAlarmChatGroup{
		BaseModel: do.BaseModel{ID: b.Id},
		Status:    b.Status,
		Remark:    b.Remark,
		Name:      b.Name,
		Hook:      b.Hook,
		NotifyApp: b.NotifyApp,
		HookName:  b.HookName,
		Template:  b.Template,
		Title:     b.Title,
		Secret:    b.Secret,
	}
}

// ChatGroupApiToBO ...
func ChatGroupApiToBO(a *api.ChatGroup) *ChatGroupBO {
	if a == nil {
		return nil
	}
	return &ChatGroupBO{
		Id:        a.Id,
		Name:      a.Name,
		Status:    vo.Status(a.Status),
		Remark:    a.Remark,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: 0,
		Hook:      a.Hook,
		NotifyApp: vo.NotifyApp(a.App),
		HookName:  a.HookName,
		Template:  a.Template,
		Title:     a.Title,
		Secret:    a.Secret,
	}
}

// ChatGroupModelToBO .
func ChatGroupModelToBO(m *do.PromAlarmChatGroup) *ChatGroupBO {
	if m == nil {
		return nil
	}
	return &ChatGroupBO{
		Id:        m.ID,
		Name:      m.Name,
		Status:    m.Status,
		Remark:    m.Remark,
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
		DeletedAt: int64(m.DeletedAt),
		Hook:      m.Hook,
		NotifyApp: m.NotifyApp,
		HookName:  m.HookName,
		Template:  m.Template,
		Title:     m.Title,
		Secret:    m.Secret,
	}
}
