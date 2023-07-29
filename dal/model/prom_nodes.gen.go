// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePromNode = "prom_nodes"

// PromNode mapped from table <prom_nodes>
type PromNode struct {
	ID         int32          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt  int64          `gorm:"column:deleted_at;type:bigint unsigned;not null" json:"deleted_at"`
	EnName     string         `gorm:"column:en_name;type:varchar(64);not null;comment:节点英文名称" json:"en_name"`            // 节点英文名称
	ChName     string         `gorm:"column:ch_name;type:varchar(64);not null;comment:节点中文名称" json:"ch_name"`            // 节点中文名称
	Datasource string         `gorm:"column:datasource;type:varchar(1024);not null;comment:prom数据源地址" json:"datasource"` // prom数据源地址
	Remark     string         `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`                 // 备注
	NodeDirs   []*PromNodeDir `gorm:"foreignKey:NodeID" json:"node_dirs"`
}

// TableName PromNode's table name
func (*PromNode) TableName() string {
	return TableNamePromNode
}
