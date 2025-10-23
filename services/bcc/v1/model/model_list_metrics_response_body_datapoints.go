package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListMetricsResponseBodyDatapoints struct {

	// 指标时间，单位毫秒
	Timestamp int64 `json:"timestamp"`

	// 数据点
	Datapoint []ListMetricsResponseBodyDatapoint `json:"datapoint"`
}

func (o ListMetricsResponseBodyDatapoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponseBodyDatapoints struct{}"
	}

	return strings.Join([]string{"ListMetricsResponseBodyDatapoints", string(data)}, " ")
}
