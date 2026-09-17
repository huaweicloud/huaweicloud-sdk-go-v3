package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRisksResponse Response Object
type ListRisksResponse struct {

	// 指标名
	MetricCode *string `json:"metric_code,omitempty"`

	// 指标展示名称
	DisplayMetricCodes *[]string `json:"display_metric_codes,omitempty"`

	// 指标名称
	MetricNames *[]string `json:"metric_names,omitempty"`

	// 单位
	Units *[]string `json:"units,omitempty"`

	// 风险实例列表
	Items          *[]RiskInfo `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRisksResponse struct{}"
	}

	return strings.Join([]string{"ListRisksResponse", string(data)}, " ")
}
