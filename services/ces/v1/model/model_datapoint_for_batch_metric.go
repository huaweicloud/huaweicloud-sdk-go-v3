package model

import (
	"encoding/json"

	"strings"
)

//
type DatapointForBatchMetric struct {
	// 聚合周期内指标数据的最大值。

	Max *float64 `json:"max,omitempty"`
	// 聚合周期内指标数据的最小值。

	Min *float64 `json:"min,omitempty"`
	// 聚合周期内指标数据的平均值。

	Average *float64 `json:"average,omitempty"`
	// 聚合周期内指标数据的求和值。

	Sum *float64 `json:"sum,omitempty"`
	// 聚合周期内指标数据的方差。

	Variance *string `json:"variance,omitempty"`
	// 指标采集时间，UNIX时间戳，单位毫秒。

	Timestamp int64 `json:"timestamp"`
}

func (o DatapointForBatchMetric) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DatapointForBatchMetric struct{}"
	}

	return strings.Join([]string{"DatapointForBatchMetric", string(data)}, " ")
}
