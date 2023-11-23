package valueobj

import (
	"prometheus-manager/api"
)

type AlarmStatus int32

const (
	// AlarmStatusUnknown 未知
	AlarmStatusUnknown AlarmStatus = iota
	// AlarmStatusAlarm 告警
	AlarmStatusAlarm
	// AlarmStatusResolved 已恢复
	AlarmStatusResolved
	// AlarmStatusIgnored 已忽略
	AlarmStatusIgnored
)

// String 转换为字符串
func (s AlarmStatus) String() string {
	switch s {
	case AlarmStatusAlarm:
		return "告警"
	case AlarmStatusResolved:
		return "已恢复"
	case AlarmStatusIgnored:
		return "已忽略"
	default:
		return "未知"
	}
}

// Key 转换为key
func (s AlarmStatus) Key() string {
	switch s {
	case AlarmStatusAlarm:
		return "alarm"
	case AlarmStatusResolved:
		return "resolved"
	case AlarmStatusIgnored:
		return "ignored"
	default:
		return "unknown"
	}
}

// Value 转换为value
func (s AlarmStatus) Value() int32 {
	return int32(s)
}

// ApiAlarmStatus 转换为api.AlarmStatus
func (s AlarmStatus) ApiAlarmStatus() api.AlarmStatus {
	return api.AlarmStatus(s)
}
