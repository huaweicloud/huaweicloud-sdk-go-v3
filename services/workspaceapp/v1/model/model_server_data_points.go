package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerDataPoints 指标数据结构。
type ServerDataPoints struct {

	// 聚合周期内指标数据的平均值。
	Average *float64 `json:"average,omitempty"`

	// 聚合周期内指标数据的最大值。
	Max *float64 `json:"max,omitempty"`

	// 聚合周期内指标数据的最小值。
	Min *float64 `json:"min,omitempty"`

	// 聚合周期内指标数据的求和值。
	Sum *float64 `json:"sum,omitempty"`

	// 聚合周期内指标数据的方差。
	Variance *float64 `json:"variance,omitempty"`

	// 指标采集时间，UNIX时间戳，单位毫秒。
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 指标单位。
	Unit *string `json:"unit,omitempty"`
}

func (o ServerDataPoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerDataPoints struct{}"
	}

	return strings.Join([]string{"ServerDataPoints", string(data)}, " ")
}
