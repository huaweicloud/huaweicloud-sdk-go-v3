package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveAlarmWhiteListRequestInfo 删除告警白名单列表
type RemoveAlarmWhiteListRequestInfo struct {

	// 删除告警白名单详情
	DataList *[]AlarmWhiteListRequestInfo `json:"data_list,omitempty"`

	// 是否需要恢复相关告警,默认 false
	RestoreAlarm *bool `json:"restore_alarm,omitempty"`

	// 是否删除所有白名单内容
	DeleteAll *bool `json:"delete_all,omitempty"`

	// 事件类型
	EventType *int32 `json:"event_type,omitempty"`
}

func (o RemoveAlarmWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAlarmWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"RemoveAlarmWhiteListRequestInfo", string(data)}, " ")
}
