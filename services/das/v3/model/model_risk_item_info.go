package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskItemInfo RiskItemInfo对象
type RiskItemInfo struct {

	// 指标码
	MetricCode *string `json:"metric_code,omitempty"`

	// 阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`
}

func (o RiskItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskItemInfo struct{}"
	}

	return strings.Join([]string{"RiskItemInfo", string(data)}, " ")
}
