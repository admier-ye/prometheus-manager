package basescopes

import (
	"strings"

	"gorm.io/gorm"
	"prometheus-manager/pkg/util/slices"
)

const (
	PreloadKeyAlarmPages        = "AlarmPages"
	PreloadKeyCategories        = "Categories"
	PreloadKeyEndpoint          = "Endpoint"
	PreloadKeyAlertLevel        = "AlertLevel"
	PreloadKeyPromNotifies      = "PromNotifies"
	PreloadKeyPromNotifyUpgrade = "PromNotifyUpgrade"
	PreloadKeyGroupInfo         = "GroupInfo"
)

const (
	StrategyTableFieldGroupID Field = "group_id"
	StrategyTableFieldAlert   Field = "alert"
)

const (
	TableNamePromStrategyAlarmPageFieldPromAlarmPageID Field = "alarm_page_id"
	TableNamePromStrategyAlarmPageFieldPromStrategyID  Field = "prom_strategy_id"
)

// StrategyTableGroupIdsEQ 策略组ID
func StrategyTableGroupIdsEQ(ids ...uint32) ScopeMethod {
	// 过滤0值
	newIds := slices.Filter(ids, func(id uint32) bool { return id > 0 })
	return WhereInColumn(StrategyTableFieldGroupID, newIds...)
}

// StrategyTableAlertLike 策略名称匹配
func StrategyTableAlertLike(keyword string) ScopeMethod {
	return WhereLikePrefixKeyword(keyword, StrategyTableFieldAlert)
}

// StrategyTablePreloadEndpoint 预加载endpoint
func StrategyTablePreloadEndpoint(db *gorm.DB) *gorm.DB {
	return db.Preload(PreloadKeyEndpoint)
}

// StrategyTablePreloadAlarmPages 预加载alarm_pages
func StrategyTablePreloadAlarmPages(db *gorm.DB) *gorm.DB {
	return db.Preload(PreloadKeyAlarmPages)
}

// StrategyTablePreloadCategories 预加载categories
func StrategyTablePreloadCategories(db *gorm.DB) *gorm.DB {
	return db.Preload(PreloadKeyCategories)
}

// StrategyTablePreloadAlertLevel 预加载alert_level
func StrategyTablePreloadAlertLevel(db *gorm.DB) *gorm.DB {
	return db.Preload(PreloadKeyAlertLevel)
}

// StrategyTablePreloadPromNotifies 预加载prom_notifies
func StrategyTablePreloadPromNotifies(preloadKeys ...string) ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		if len(preloadKeys) == 0 {
			return db.Preload(PreloadKeyPromNotifies)
		}
		tx := db
		for _, key := range preloadKeys {
			tx = tx.Preload(strings.Join([]string{PreloadKeyPromNotifies, key}, "."))
		}
		return tx
	}
}

// StrategyTablePreloadPromNotifyUpgrade 预加载prom_notify_upgrade
func StrategyTablePreloadPromNotifyUpgrade(db *gorm.DB) *gorm.DB {
	return db.Preload(PreloadKeyPromNotifyUpgrade)
}

// StrategyTablePreloadGroupInfo 预加载group_info
func StrategyTablePreloadGroupInfo(db *gorm.DB) *gorm.DB {
	return db.Preload(PreloadKeyGroupInfo)
}

// InTableNamePromStrategyAlarmPageFieldPromAlarmPageIds in alarm_page_ids
func InTableNamePromStrategyAlarmPageFieldPromAlarmPageIds(ids ...uint32) ScopeMethod {
	return WhereInColumn(TableNamePromStrategyAlarmPageFieldPromAlarmPageID, ids...)
}

// InTableNamePromStrategyAlarmPageFieldPromStrategyID in prom_strategy_ids
func InTableNamePromStrategyAlarmPageFieldPromStrategyID(ids ...uint32) ScopeMethod {
	return WhereInColumn(TableNamePromStrategyAlarmPageFieldPromStrategyID, ids...)
}
