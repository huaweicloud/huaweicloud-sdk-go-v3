package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Datapoint 时间点的统计数据
type Datapoint struct {

	// 最大值，未计算默认为0
	Max *float64 `json:"max,omitempty"`

	// 最小值，未计算默认为0
	Min *float64 `json:"min,omitempty"`

	// 平均值，未计算默认为0
	Average *float64 `json:"average,omitempty"`

	// 综合，未计算默认为0
	Sum *float64 `json:"sum,omitempty"`

	// 方差，未计算默认为0
	Variance *float64 `json:"variance,omitempty"`

	// 毫秒时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 数据单位，比如%，个
	Unit *string `json:"unit,omitempty"`
}

func (o Datapoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datapoint struct{}"
	}

	return strings.Join([]string{"Datapoint", string(data)}, " ")
}
