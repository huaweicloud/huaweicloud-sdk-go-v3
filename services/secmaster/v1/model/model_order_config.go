package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderConfig 订单配置项
type OrderConfig struct {

	// 阈值列表
	ThresholdList *[]UsageThreshold `json:"threshold_list,omitempty"`

	AlertConfig *AlertConfig `json:"alert_config,omitempty"`
}

func (o OrderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderConfig struct{}"
	}

	return strings.Join([]string{"OrderConfig", string(data)}, " ")
}
