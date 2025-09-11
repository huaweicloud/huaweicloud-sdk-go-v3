package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDataPoint **参数解释**： 监控数据点
type MetricDataPoint struct {

	// **参数解释**： 维度信息
	Dimensions *[]MetricDataPointDimensions `json:"dimensions,omitempty"`

	// **参数解释**： 指标采集时间 **取值范围**： 最小值为0
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释**： 指标值 **取值范围**： 不涉及
	Value *float64 `json:"value,omitempty"`

	// **参数解释**： 数据的单位。    **取值范围**： 长度为[0,32]个字符。
	Unit *string `json:"unit,omitempty"`
}

func (o MetricDataPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataPoint struct{}"
	}

	return strings.Join([]string{"MetricDataPoint", string(data)}, " ")
}
