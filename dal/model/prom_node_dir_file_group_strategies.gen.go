// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePromNodeDirFileGroupStrategy = "prom_node_dir_file_group_strategies"

// PromNodeDirFileGroupStrategy mapped from table <prom_node_dir_file_group_strategies>
type PromNodeDirFileGroupStrategy struct {
	ID          int32        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	GroupID     int32        `gorm:"column:group_id;type:int unsigned;not null;comment:策略组 ID" json:"group_id"`       // 策略组 ID
	Alert       string       `gorm:"column:alert;type:varchar(64);not null;default:_rules;comment:策略名称" json:"alert"` // 策略名称
	Expr        string       `gorm:"column:expr;type:varchar(4096);not null;default:0;comment:prom QL" json:"expr"`   // prom QL
	For         *string      `gorm:"column:for;type:varchar(32);default:7200h;comment:持续时间" json:"for"`               // 持续时间
	Labels      string       `gorm:"column:labels;type:json;not null;comment:标签" json:"labels"`                       // 标签
	Annotations string       `gorm:"column:annotations;type:json;not null;comment:内容" json:"annotations"`             // 内容
	CreatedAt   time.Time    `gorm:"column:created_at;type:timestamp;not null;comment:创建时间" json:"created_at"`        // 创建时间
	UpdatedAt   *time.Time   `gorm:"column:updated_at;type:timestamp;comment:更新时间" json:"updated_at"`                 // 更新时间
	DeletedAt   int64        `gorm:"column:deleted_at;type:bigint;not null;comment:删除时间" json:"deleted_at"`           // 删除时间
	Strategies  []*PromCombo `gorm:"foreignKey:strategy_id" json:"strategies"`
}

// TableName PromNodeDirFileGroupStrategy's table name
func (*PromNodeDirFileGroupStrategy) TableName() string {
	return TableNamePromNodeDirFileGroupStrategy
}
