package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateBaselineRequest struct {

	// 工作空间ID。
	WorkspaceId string `json:"workspace_id"`

	// 基线任务名称。只能包含：英文字母、数字、中文、中划线、下划线和点号，且长度不超过128位。
	Name string `json:"name"`

	// 责任人用户名称。
	OwnerName string `json:"owner_name"`

	// 基线任务类型。取值为 DAY和HOUR。
	Type string `json:"type"`

	// 保障作业ID列表。
	SlaTaskIds []string `json:"sla_task_ids"`

	// 优先级，取值范围是1到5。
	Priority int32 `json:"priority"`

	// 预警余量。单位毫秒，取值范围在0到86400000之间，且必须是整数分钟级别的毫秒。
	Buffer int32 `json:"buffer"`

	// 是否生效，取值为true或者false。默认为true。
	Enable bool `json:"enable"`

	// 报警是否打开，取值为true或者false，默认为true。
	AlarmEnable bool `json:"alarm_enable"`

	// 基线签署是否打开，取值为true或者false。默认为false。
	SignEnable bool `json:"sign_enable"`

	// 天基线承诺小时。
	SlaHour string `json:"sla_hour"`

	// 天基线承诺分钟。
	SlaMin string `json:"sla_min"`

	// 基线报警是否打开，取值为true或者false。
	BaselineAlarmEnable bool `json:"baseline_alarm_enable"`

	// SMN主题列表。
	SmnTopics *[]SmnTopicRequest `json:"smn_topics,omitempty"`

	// 事件报警信息。
	EventSmnTopics *[]SmnTopicRequest `json:"event_smn_topics,omitempty"`

	// 事件告警开启类型，取值为ERROR：出错，SLOW_DOWN：变慢。
	EventAlarm *[]string `json:"event_alarm,omitempty"`

	// 当type取值为HOUR时，该值需要填写。
	Period *[]PeriodObject `json:"period,omitempty"`
}

func (o CreateBaselineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBaselineRequest struct{}"
	}

	return strings.Join([]string{"CreateBaselineRequest", string(data)}, " ")
}
