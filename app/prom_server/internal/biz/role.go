package biz

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	"github.com/go-kratos/kratos/v2/log"
	"prometheus-manager/api"
	"prometheus-manager/app/prom_server/internal/biz/dobo"
	"prometheus-manager/app/prom_server/internal/biz/repository"
	"prometheus-manager/pkg/helper/model/system"
)

type (
	RoleBiz struct {
		log *log.Helper

		roleRepo repository.RoleRepo
	}
)

func NewRoleBiz(roleRepo repository.RoleRepo, logger log.Logger) *RoleBiz {
	return &RoleBiz{
		log:      log.NewHelper(logger),
		roleRepo: roleRepo,
	}
}

// CreateRole 创建角色
func (b *RoleBiz) CreateRole(ctx context.Context, role *dobo.RoleBO) (*dobo.RoleBO, error) {
	do := dobo.NewRoleBO(role).DO().First()
	do, err := b.roleRepo.Create(ctx, do)
	if err != nil {
		return nil, err
	}

	return dobo.NewRoleDO(do).BO().First(), nil
}

// DeleteRoleByIds 删除角色
func (b *RoleBiz) DeleteRoleByIds(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return b.roleRepo.Delete(ctx, system.RoleInIds(ids...))
}

// ListRole 角色列表
func (b *RoleBiz) ListRole(ctx context.Context, pgInfo query.Pagination, scopes ...query.ScopeMethod) ([]*dobo.RoleBO, error) {
	dList, err := b.roleRepo.List(ctx, pgInfo, scopes...)
	if err != nil {
		return nil, err
	}

	return dobo.NewRoleDO(dList...).BO().List(), nil
}

// GetRoleById 获取角色
func (b *RoleBiz) GetRoleById(ctx context.Context, id uint32) (*dobo.RoleBO, error) {
	do, err := b.roleRepo.Get(ctx, system.RoleInIds(id))
	if err != nil {
		return nil, err
	}

	return dobo.NewRoleDO(do).BO().First(), nil
}

// UpdateRoleById 更新角色
func (b *RoleBiz) UpdateRoleById(ctx context.Context, role *dobo.RoleBO) (*dobo.RoleBO, error) {
	do := dobo.NewRoleBO(role).DO().First()
	do, err := b.roleRepo.Update(ctx, do, system.RoleInIds(role.Id))
	if err != nil {
		return nil, err
	}

	return dobo.NewRoleDO(do).BO().First(), nil
}

// UpdateRoleStatusById 更新角色状态
func (b *RoleBiz) UpdateRoleStatusById(ctx context.Context, status api.Status, ids []uint32) error {
	do := &dobo.RoleDO{Status: int32(status)}
	do, err := b.roleRepo.Update(ctx, do, system.RoleInIds(ids...))
	if err != nil {
		return err
	}

	return nil
}
