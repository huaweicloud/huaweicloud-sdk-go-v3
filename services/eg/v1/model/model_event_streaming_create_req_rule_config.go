package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventStreamingCreateReqRuleConfig 事件规则，包括过滤规则和转换规则
type EventStreamingCreateReqRuleConfig struct {
	Transform *TransForm `json:"transform,omitempty"`

	// 过滤规则
	Filter *interface{} `json:"filter,omitempty"`
}

func (o EventStreamingCreateReqRuleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventStreamingCreateReqRuleConfig struct{}"
	}

	return strings.Join([]string{"EventStreamingCreateReqRuleConfig", string(data)}, " ")
}
