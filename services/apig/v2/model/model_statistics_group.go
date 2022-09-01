package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsGroup struct {

	// 最大延时  单位：ms
	MaxLatency *int32 `json:"max_latency,omitempty" xml:"max_latency"`

	// 平均延时  单位：ms
	AvgLatency *float32 `json:"avg_latency,omitempty" xml:"avg_latency"`

	// 请求总次数
	ReqCount *int32 `json:"req_count,omitempty" xml:"req_count"`

	// 2xx响应码总次数
	ReqCount2xx *int32 `json:"req_count2xx,omitempty" xml:"req_count2xx"`

	// 4xx响应码总次数
	ReqCount4xx *int32 `json:"req_count4xx,omitempty" xml:"req_count4xx"`

	// 5xx响应码总次数
	ReqCount5xx *int32 `json:"req_count5xx,omitempty" xml:"req_count5xx"`

	// 错误次数
	ReqCountError *int32 `json:"req_count_error,omitempty" xml:"req_count_error"`

	// 下行吞吐量（byte）
	OutputThroughput *int64 `json:"output_throughput,omitempty" xml:"output_throughput"`

	// 上行吞吐量（byte）
	InputThroughput *int64 `json:"input_throughput,omitempty" xml:"input_throughput"`

	// API访问的UTC时间戳
	CurrentMinute *int64 `json:"current_minute,omitempty" xml:"current_minute"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// API拥有者
	Provider *string `json:"provider,omitempty" xml:"provider"`

	// API请求时间
	ReqTime *sdktime.SdkTime `json:"req_time,omitempty" xml:"req_time"`

	// 记录时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty" xml:"register_time"`
}

func (o StatisticsGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsGroup struct{}"
	}

	return strings.Join([]string{"StatisticsGroup", string(data)}, " ")
}
