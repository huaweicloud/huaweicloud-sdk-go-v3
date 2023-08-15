package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricData struct {

	// 聚合周期内指标数据的平均值，仅当请求参数filter字段值为average时支持。
	Average *int32 `json:"average,omitempty"`

	// 聚合周期内指标数据的最大值，仅当请求参数filter字段值为max时支持。
	Max *int32 `json:"max,omitempty"`

	// 聚合周期内指标数据的最小值，仅当请求参数filter字段值为min时支持。
	Min *int32 `json:"min,omitempty"`

	// 聚合周期内指标数据的求和值，仅当请求参数filter字段值为sum时支持。
	Sum *int32 `json:"sum,omitempty"`

	// 聚合周期内指标数据的方差，仅当请求参数filter字段值为variance时支持。
	Variance *int32 `json:"variance,omitempty"`

	// 指标采集时间，UNIX时间戳，单位毫秒。
	Timestamp *int32 `json:"timestamp,omitempty"`

	// 指标单位。
	Unit *string `json:"unit,omitempty"`
}

func (o MetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricData struct{}"
	}

	return strings.Join([]string{"MetricData", string(data)}, " ")
}
