package do

import (
	"encoding/json"

	"prometheus-manager/app/prom_server/internal/biz/vo"

	"prometheus-manager/pkg/strategy"
)

const TableNamePromStrategy = "prom_strategies"

// PromStrategy mapped from table <prom_strategies>
type PromStrategy struct {
	BaseModel
	GroupID      uint32                `gorm:"column:group_id;type:int unsigned;not null;comment:所属规则组ID" json:"group_id"`
	Alert        string                `gorm:"column:alert;type:varchar(64);not null;comment:规则名称" json:"alert"`
	Expr         string                `gorm:"column:expr;type:text;not null;comment:prom ql" json:"expr"`
	For          string                `gorm:"column:for;type:varchar(64);not null;default:10s;comment:持续时间" json:"for"`
	Labels       *strategy.Labels      `gorm:"column:labels;type:json;not null;comment:标签" json:"labels"`
	Annotations  *strategy.Annotations `gorm:"column:annotations;type:json;not null;comment:告警文案" json:"annotations"`
	AlertLevelID uint32                `gorm:"column:alert_level_id;type:int;not null;index:idx__alert_level_id,priority:1;comment:告警等级dict ID" json:"alert_level_id"`
	Status       vo.Status             `gorm:"column:status;type:tinyint;not null;default:1;comment:启用状态: 1启用;2禁用" json:"status"`
	Remark       string                `gorm:"column:remark;type:varchar(255);not null;comment:描述信息" json:"remark"`

	AlarmPages []*PromAlarmPage   `gorm:"References:ID;foreignKey:ID;joinForeignKey:PromStrategyID;joinReferences:AlarmPageID;many2many:prom_strategy_alarm_pages" json:"-"`
	Categories []*PromDict        `gorm:"References:ID;foreignKey:ID;joinForeignKey:PromStrategyID;joinReferences:DictID;many2many:prom_strategy_categories" json:"-"`
	AlertLevel *PromDict          `gorm:"foreignKey:AlertLevelID" json:"-"`
	GroupInfo  *PromStrategyGroup `gorm:"foreignKey:GroupID" json:"-"`

	// 通知对象
	PromNotifies []*PromAlarmNotify `gorm:"many2many:prom_strategy_notifies;comment:通知对象" json:"-"`
	// 告警升级后的通知对象
	PromNotifyUpgrade []*PromAlarmNotify `gorm:"many2many:prom_strategy_notify_upgrades;comment:告警升级后的通知对象" json:"-"`
	// 最大抑制时长(s)
	MaxSuppress string `gorm:"column:max_suppress;type:varchar(255);not null;default:1m;comment:最大抑制时长(s)" json:"max_suppress"`
	// 是否发送告警恢复通知
	SendRecover vo.IsSendRecover `gorm:"column:send_recover;type:tinyint;not null;default:0;comment:是否发送告警恢复通知" json:"send_recover"`
	// 发送告警时间间隔(s), 默认为for的10倍时间分钟, 用于长时间未消警情况
	SendInterval string `gorm:"column:send_interval;type:varchar(255);not null;default:1m;comment:发送告警时间间隔(s), 默认为for的10倍时间分钟, 用于长时间未消警情况" json:"send_interval"`

	// 数据源
	EndpointID uint32    `gorm:"column:endpoint_id;type:int unsigned;not null;default:0;comment:数据源ID" json:"endpoint_id"`
	Endpoint   *Endpoint `gorm:"foreignKey:EndpointID" json:"-"`

	// 创建人ID
	CreateBy uint32 `gorm:"column:create_by;type:int;not null;comment:创建人ID"`
}

// TableName PromStrategy's table name
func (*PromStrategy) TableName() string {
	return TableNamePromStrategy
}

// GetAlertLevel 获取告警等级
func (p *PromStrategy) GetAlertLevel() *PromDict {
	if p == nil {
		return nil
	}
	return p.AlertLevel
}

// GetAlarmPages 获取告警页面
func (p *PromStrategy) GetAlarmPages() []*PromAlarmPage {
	if p == nil {
		return nil
	}
	return p.AlarmPages
}

// GetLabels 获取标签
func (p *PromStrategy) GetLabels() *strategy.Labels {
	if p == nil {
		return nil
	}
	return p.Labels
}

// GetAnnotations 获取告警文案
func (p *PromStrategy) GetAnnotations() *strategy.Annotations {
	if p == nil {
		return nil
	}
	return p.Annotations
}

// GetCategories 获取分类
func (p *PromStrategy) GetCategories() []*PromDict {
	if p == nil {
		return nil
	}
	return p.Categories
}

// GetPromNotifies 获取通知对象
func (p *PromStrategy) GetPromNotifies() []*PromAlarmNotify {
	if p == nil {
		return nil
	}
	return p.PromNotifies
}

// GetPromNotifyUpgrade 获取告警升级后的通知对象
func (p *PromStrategy) GetPromNotifyUpgrade() []*PromAlarmNotify {
	if p == nil {
		return nil
	}
	return p.PromNotifyUpgrade
}

// GetGroupInfo 获取所属规则组
func (p *PromStrategy) GetGroupInfo() *PromStrategyGroup {
	if p == nil {
		return nil
	}
	return p.GroupInfo
}

// GetEndpoint 获取数据源
func (p *PromStrategy) GetEndpoint() *Endpoint {
	if p == nil {
		return nil
	}
	return p.Endpoint
}

// ToMap to map[string]any
func (p *PromStrategy) ToMap() map[string]any {
	if p == nil {
		return nil
	}
	bs, _ := json.Marshal(p)
	pMap := make(map[string]any)
	_ = json.Unmarshal(bs, &pMap)
	pMap["labels"] = p.Labels.String()
	pMap["annotations"] = p.Annotations.String()
	return pMap
}
