package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Datapoint struct {

	// 聚合周期内指标数据的最大值。
	Max *float64 `json:"max,omitempty" xml:"max"`

	// 聚合周期内指标数据的最小值。
	Min *float64 `json:"min,omitempty" xml:"min"`

	// 聚合周期内指标数据的平均值。
	Average *float64 `json:"average,omitempty" xml:"average"`

	// 聚合周期内指标数据的求和值。
	Sum *float64 `json:"sum,omitempty" xml:"sum"`

	// 聚合周期内指标数据的方差。
	Variance *float64 `json:"variance,omitempty" xml:"variance"`

	// 指标采集时间，UNIX时间戳，单位毫秒。
	Timestamp int64 `json:"timestamp" xml:"timestamp"`

	// 指标单位。
	Unit *string `json:"unit,omitempty" xml:"unit"`
}

func (o Datapoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datapoint struct{}"
	}

	return strings.Join([]string{"Datapoint", string(data)}, " ")
}
