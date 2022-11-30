package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Event2alarmRuleBodyTriggerPolicies struct {

	// 自增编号
	Id *int32 `json:"id,omitempty"`

	// 事件名称
	Name *string `json:"name,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 触发周期
	Period *int32 `json:"period,omitempty"`

	// 比较符
	Operator *string `json:"operator,omitempty"`

	// 触发次数
	Count *int32 `json:"count,omitempty"`

	// 告警等级
	Level *string `json:"level,omitempty"`
}

func (o Event2alarmRuleBodyTriggerPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event2alarmRuleBodyTriggerPolicies struct{}"
	}

	return strings.Join([]string{"Event2alarmRuleBodyTriggerPolicies", string(data)}, " ")
}
