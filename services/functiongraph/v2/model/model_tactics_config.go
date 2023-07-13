package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TacticsConfig struct {

	// 定时配置列表
	CronConfigs *[]CronConfig `json:"cron_configs,omitempty"`

	// 流量配置列表
	MetricConfigs *[]MetricConfig `json:"metric_configs,omitempty"`
}

func (o TacticsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TacticsConfig struct{}"
	}

	return strings.Join([]string{"TacticsConfig", string(data)}, " ")
}
