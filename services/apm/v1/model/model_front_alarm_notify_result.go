package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrontAlarmNotifyResult 告警通知内容。
type FrontAlarmNotifyResult struct {

	// 告警通知id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 通知类型。
	NotifyType *string `json:"notify_type,omitempty"`

	// 告警规则id。
	AlarmRuleId *int64 `json:"alarm_rule_id,omitempty"`

	// 模板id。
	TemplateId *int64 `json:"template_id,omitempty"`

	// 关联事件id。
	AlarmDataEventId *int64 `json:"alarm_data_event_id,omitempty"`

	// 通知结果。
	NotifyStatus *bool `json:"notify_status,omitempty"`

	// 通知内容。
	AlarmContent *string `json:"alarm_content,omitempty"`
}

func (o FrontAlarmNotifyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontAlarmNotifyResult struct{}"
	}

	return strings.Join([]string{"FrontAlarmNotifyResult", string(data)}, " ")
}
