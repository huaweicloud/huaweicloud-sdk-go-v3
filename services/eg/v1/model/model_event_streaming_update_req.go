package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventStreamingUpdateReq struct {

	// 事件流名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须字母或数字开头
	Name string `json:"name"`

	// 事件流描述
	Description *string `json:"description,omitempty"`

	Source *EventStreamingSource `json:"source"`

	Sink *EventStreamingSink `json:"sink"`

	RuleConfig *EventStreamingCreateReqRuleConfig `json:"rule_config,omitempty"`

	Option *RunOption `json:"option,omitempty"`
}

func (o EventStreamingUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventStreamingUpdateReq struct{}"
	}

	return strings.Join([]string{"EventStreamingUpdateReq", string(data)}, " ")
}
