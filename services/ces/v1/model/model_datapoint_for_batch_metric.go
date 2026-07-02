package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatapointForBatchMetric
type DatapointForBatchMetric struct {

	// **参数解释** 聚合周期内指标数据的最大值 **取值范围** 不涉及
	Max *float64 `json:"max,omitempty"`

	// **参数解释** 聚合周期内指标数据的最小值 **取值范围** 不涉及
	Min *float64 `json:"min,omitempty"`

	// **参数解释** 聚合周期内指标数据的平均值 **取值范围** 不涉及
	Average *float64 `json:"average,omitempty"`

	// **参数解释** 聚合周期内指标数据的求和值 **取值范围** 不涉及
	Sum *float64 `json:"sum,omitempty"`

	// **参数解释** 聚合周期内指标数据的方差 **取值范围** 不涉及
	Variance *float64 `json:"variance,omitempty"`

	// **参数解释** 指标采集时间，UNIX时间戳，单位毫秒 **取值范围** 不涉及
	Timestamp int64 `json:"timestamp"`
}

func (o DatapointForBatchMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatapointForBatchMetric struct{}"
	}

	return strings.Join([]string{"DatapointForBatchMetric", string(data)}, " ")
}
