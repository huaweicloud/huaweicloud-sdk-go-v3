package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListMetricDataRequestBody
type BatchListMetricDataRequestBody struct {

	// 指标数据。数组长度最大500
	Metrics []MetricInfo `json:"metrics"`

	Period *BatchPeriod `json:"period"`

	Filter *Filter `json:"filter"`

	From int64 `json:"from"`

	To int64 `json:"to"`
}

func (o BatchListMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequestBody", string(data)}, " ")
}
