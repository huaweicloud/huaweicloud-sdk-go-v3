package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HighFrequencyScanBlockingConfig struct {
	Action *PolicyAction `json:"action,omitempty"`

	// 检测时间范围，单位秒
	TimeRange *int32 `json:"time_range,omitempty"`

	// 触发阈值，在检测时间内扫描次数超过该值则触发封禁
	TriggerThreshold *int32 `json:"trigger_threshold,omitempty"`

	// 最大触发规则数，在检测时间内命中不同防护规则的数量达到该值时触发封禁
	MaxNumTriggeredRules *int32 `json:"max_num_triggered_rules,omitempty"`

	// 封禁周期，单位秒
	BlockingPeriod *int32 `json:"blocking_period,omitempty"`
}

func (o HighFrequencyScanBlockingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HighFrequencyScanBlockingConfig struct{}"
	}

	return strings.Join([]string{"HighFrequencyScanBlockingConfig", string(data)}, " ")
}
