package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendPoint struct {

	// 时间点（毫秒时间戳）
	Timestamp *int64 `json:"timestamp,omitempty"`

	// key大小（size/storage维度返回）
	Size *int64 `json:"size,omitempty"`

	// QPS值（qps维度返回）
	Qps *int64 `json:"qps,omitempty"`

	// 入口流量（traffic维度返回，字节/秒）
	TrafficIn *int64 `json:"traffic_in,omitempty"`

	// 出口流量（traffic维度返回，字节/秒）
	TrafficOut *int64 `json:"traffic_out,omitempty"`
}

func (o TrendPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendPoint struct{}"
	}

	return strings.Join([]string{"TrendPoint", string(data)}, " ")
}
