package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedMetricNamesResponse Response Object
type ListSupportedMetricNamesResponse struct {

	// 支持指标名称列表
	Items          *[]MetricNamesSupportItem `json:"items,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListSupportedMetricNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedMetricNamesResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedMetricNamesResponse", string(data)}, " ")
}
