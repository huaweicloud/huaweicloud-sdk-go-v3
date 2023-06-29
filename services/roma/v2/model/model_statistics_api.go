package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type StatisticsApi struct {

	// 最大延时
	MaxLatency *int32 `json:"max_latency,omitempty"`

	// 平均延时
	AvgLatency *float32 `json:"avg_latency,omitempty"`

	// 请求总次数
	ReqCount *int32 `json:"req_count,omitempty"`

	// 2xx响应码总次数
	ReqCount2xx *int32 `json:"req_count2xx,omitempty"`

	// 4xx响应码总次数
	ReqCount4xx *int32 `json:"req_count4xx,omitempty"`

	// 5xx响应码总次数
	ReqCount5xx *int32 `json:"req_count5xx,omitempty"`

	// 错误次数
	ReqCountError *int32 `json:"req_count_error,omitempty"`

	// 最大网关内部延时
	MaxInnerLatency *int32 `json:"max_inner_latency,omitempty"`

	// 平均网关内部延时
	AvgInnerLatency *float32 `json:"avg_inner_latency,omitempty"`

	// 最大后端延时
	MaxBackendLatency *int32 `json:"max_backend_latency,omitempty"`

	// 平均后端延时
	AvgBackendLatency *float32 `json:"avg_backend_latency,omitempty"`

	// 下行吞吐量（byte）
	OutputThroughput *int64 `json:"output_throughput,omitempty"`

	// 上行吞吐量（byte）
	InputThroughput *int64 `json:"input_throughput,omitempty"`

	// API访问的UTC时间戳
	CurrentMinute *int64 `json:"current_minute,omitempty"`

	// 查询统计周期
	Cycle *StatisticsApiCycle `json:"cycle,omitempty"`

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// API拥有者
	Provider *string `json:"provider,omitempty"`

	// API请求时间
	ReqTime *sdktime.SdkTime `json:"req_time,omitempty"`

	// 记录时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 状态值： - 1：调度中，未上报CES - 2：已经成功上报CES  预留字段，暂不支持
	Status *StatisticsApiStatus `json:"status,omitempty"`
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type StatisticsApiStatus struct {
	value int32
}

type StatisticsApiStatusEnum struct {
	E_1 StatisticsApiStatus
	E_2 StatisticsApiStatus
}

func GetStatisticsApiStatusEnum() StatisticsApiStatusEnum {
	return StatisticsApiStatusEnum{
		E_1: StatisticsApiStatus{
			value: 1,
		}, E_2: StatisticsApiStatus{
			value: 2,
		},
	}
}

func (c StatisticsApiStatus) Value() int32 {
	return c.value
}

func (c StatisticsApiStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatisticsApiStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
