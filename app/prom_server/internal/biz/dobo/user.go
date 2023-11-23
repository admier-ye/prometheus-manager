package dobo

import (
	"time"

	"prometheus-manager/api"
	"prometheus-manager/app/prom_server/internal/biz/valueobj"
)

type (
	UserDO struct {
		Id        uint      `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Status    int32     `json:"status"`
		Remark    string    `json:"remark"`
		Avatar    string    `json:"avatar"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		DeletedAt int64     `json:"deletedAt"`
	}

	UserBO struct {
		Id        uint            `json:"id"`
		Username  string          `json:"username"`
		Password  string          `json:"password"`
		Email     string          `json:"email"`
		Phone     string          `json:"phone"`
		Status    valueobj.Status `json:"status"`
		Remark    string          `json:"remark"`
		Avatar    string          `json:"avatar"`
		CreatedAt int64           `json:"createdAt"`
		UpdatedAt int64           `json:"updatedAt"`
		DeletedAt int64           `json:"deletedAt"`
	}
)

// NewUserDO .
func NewUserDO(values ...*UserDO) IDO[*UserBO, *UserDO] {
	return NewDO[*UserBO, *UserDO](
		DOWithValues[*UserBO, *UserDO](values...),
		DOWithBToD[*UserBO, *UserDO](userBoToDo),
		DOWithDToB[*UserBO, *UserDO](userDoToBo),
	)
}

// NewUserBO .
func NewUserBO(values ...*UserBO) IBO[*UserBO, *UserDO] {
	return NewBO[*UserBO, *UserDO](
		BOWithValues[*UserBO, *UserDO](values...),
		BOWithDToB[*UserBO, *UserDO](userDoToBo),
		BOWithBToD[*UserBO, *UserDO](userBoToDo),
	)
}

func userDoToBo(d *UserDO) *UserBO {
	if d == nil {
		return nil
	}
	return &UserBO{
		Id:        d.Id,
		Username:  d.Username,
		Password:  d.Password,
		Email:     d.Email,
		Phone:     d.Phone,
		Status:    valueobj.Status(d.Status),
		Remark:    d.Remark,
		Avatar:    d.Avatar,
		CreatedAt: d.CreatedAt.Unix(),
		UpdatedAt: d.UpdatedAt.Unix(),
		DeletedAt: d.DeletedAt,
	}
}

func userBoToDo(b *UserBO) *UserDO {
	if b == nil {
		return nil
	}
	return &UserDO{
		Id:        b.Id,
		Username:  b.Username,
		Password:  b.Password,
		Email:     b.Email,
		Phone:     b.Phone,
		Status:    int32(b.Status),
		Remark:    b.Remark,
		Avatar:    b.Avatar,
		CreatedAt: time.Unix(b.CreatedAt, 0),
		UpdatedAt: time.Unix(b.UpdatedAt, 0),
		DeletedAt: b.DeletedAt,
	}
}

func (l *UserBO) ToApiUserSelectV1() *api.UserSelectV1 {
	if l == nil {
		return nil
	}

	return &api.UserSelectV1{
		Value:  uint32(l.Id),
		Label:  l.Username,
		Status: l.Status.ApiStatus(),
		Remark: l.Remark,
		Avatar: l.Avatar,
	}
}

func (l *UserBO) ToApiUserV1() *api.UserV1 {
	if l == nil {
		return nil
	}

	return &api.UserV1{
		Id:        uint32(l.Id),
		Username:  l.Username,
		Email:     l.Email,
		Phone:     l.Phone,
		Status:    l.Status.ApiStatus(),
		Remark:    l.Remark,
		Avatar:    l.Avatar,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
		DeletedAt: l.DeletedAt,
	}
}
