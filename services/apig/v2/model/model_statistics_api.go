package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type StatisticsApi struct {

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

	// 最大网关内部延时  单位：ms
	MaxInnerLatency *int32 `json:"max_inner_latency,omitempty" xml:"max_inner_latency"`

	// 平均网关内部延时  单位：ms
	AvgInnerLatency *float32 `json:"avg_inner_latency,omitempty" xml:"avg_inner_latency"`

	// 最大后端延时
	MaxBackendLatency *int32 `json:"max_backend_latency,omitempty" xml:"max_backend_latency"`

	// 平均后端延时
	AvgBackendLatency *float32 `json:"avg_backend_latency,omitempty" xml:"avg_backend_latency"`

	// 下行吞吐量（byte）
	OutputThroughput *int64 `json:"output_throughput,omitempty" xml:"output_throughput"`

	// 上行吞吐量（byte）
	InputThroughput *int64 `json:"input_throughput,omitempty" xml:"input_throughput"`

	// API访问的UTC时间戳
	CurrentMinute *int64 `json:"current_minute,omitempty" xml:"current_minute"`

	// 统计周期
	Cycle *StatisticsApiCycle `json:"cycle,omitempty" xml:"cycle"`

	// API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// API拥有者
	Provider *string `json:"provider,omitempty" xml:"provider"`

	// API请求时间
	ReqTime *sdktime.SdkTime `json:"req_time,omitempty" xml:"req_time"`

	// 记录时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty" xml:"register_time"`

	// 状态
	Status *int32 `json:"status,omitempty" xml:"status"`
}

func (o StatisticsApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsApi struct{}"
	}

	return strings.Join([]string{"StatisticsApi", string(data)}, " ")
}

type StatisticsApiCycle struct {
	value string
}

type StatisticsApiCycleEnum struct {
	MINUTE StatisticsApiCycle
	HOUR   StatisticsApiCycle
	DAY    StatisticsApiCycle
}

func GetStatisticsApiCycleEnum() StatisticsApiCycleEnum {
	return StatisticsApiCycleEnum{
		MINUTE: StatisticsApiCycle{
			value: "MINUTE",
		},
		HOUR: StatisticsApiCycle{
			value: "HOUR",
		},
		DAY: StatisticsApiCycle{
			value: "DAY",
		},
	}
}

func (c StatisticsApiCycle) Value() string {
	return c.value
}

func (c StatisticsApiCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatisticsApiCycle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
