package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRiskItemsItems struct {

	// 指标码
	MetricCode *string `json:"metric_code,omitempty"`

	// 阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`
}

func (o QueryRiskItemsItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRiskItemsItems struct{}"
	}

	return strings.Join([]string{"QueryRiskItemsItems", string(data)}, " ")
}
