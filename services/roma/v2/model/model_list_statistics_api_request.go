package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListStatisticsApiRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 查询模式，默认为INSTANCE * ALL：实例下所有调用应用，要求主帐号权限 * APP：指定集成应用 * API：指定API * INSTANCE：实例，默认值  注意：mode = APP或ALL时，接口响应中不返回cycle，api_id，group_id，provider，register_time，status字段
	Mode *ListStatisticsApiRequestMode `json:"mode,omitempty"`

	// 集成应用编号，查询模式为APP时必填
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// API编号，查询模式为API时必填
	ApiId *string `json:"api_id,omitempty"`

	// 查询统计周期 * minute：分钟 * hour：小时 * day：天
	Cycle *ListStatisticsApiRequestCycle `json:"cycle,omitempty"`

	// 开始时间，格式：2020-06-18 10:00:01
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式：2020-06-18 23:00:00
	EndTime *string `json:"end_time,omitempty"`

	// 统计时长格式：整数+单位（m、h），m：分钟，h：小时，可支持小时与分钟的组合。例如：1h或2h45m * 同时给定start_time和end_time优先查询[start_time, end_time] * start_time不存在，end_time和duration存在且合法，则查询区间为[end_time - duration, end_time] * start_time和end_time不存在，duration存在且合法，令end_time=now，则查询区间为[end_time - duration, end_time] * start_time，end_time和duration都不存在，报错missing time range parameters。 * duration最长查询范围：小时最长支持72小时，分钟最长支持90分钟。
	Duration *string `json:"duration,omitempty"`
}

func (o ListStatisticsApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsApiRequest struct{}"
	}

	return strings.Join([]string{"ListStatisticsApiRequest", string(data)}, " ")
}

type ListStatisticsApiRequestMode struct {
	value string
}

type ListStatisticsApiRequestModeEnum struct {
	INSTANCE ListStatisticsApiRequestMode
	ALL      ListStatisticsApiRequestMode
	APP      ListStatisticsApiRequestMode
	API      ListStatisticsApiRequestMode
}

func GetListStatisticsApiRequestModeEnum() ListStatisticsApiRequestModeEnum {
	return ListStatisticsApiRequestModeEnum{
		INSTANCE: ListStatisticsApiRequestMode{
			value: "INSTANCE",
		},
		ALL: ListStatisticsApiRequestMode{
			value: "ALL",
		},
		APP: ListStatisticsApiRequestMode{
			value: "APP",
		},
		API: ListStatisticsApiRequestMode{
			value: "API",
		},
	}
}

func (c ListStatisticsApiRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStatisticsApiRequestMode) UnmarshalJSON(b []byte) error {
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

type ListStatisticsApiRequestCycle struct {
	value string
}

type ListStatisticsApiRequestCycleEnum struct {
	MINUTE ListStatisticsApiRequestCycle
	HOUR   ListStatisticsApiRequestCycle
	DAY    ListStatisticsApiRequestCycle
}

func GetListStatisticsApiRequestCycleEnum() ListStatisticsApiRequestCycleEnum {
	return ListStatisticsApiRequestCycleEnum{
		MINUTE: ListStatisticsApiRequestCycle{
			value: "minute",
		},
		HOUR: ListStatisticsApiRequestCycle{
			value: "hour",
		},
		DAY: ListStatisticsApiRequestCycle{
			value: "day",
		},
	}
}

func (c ListStatisticsApiRequestCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStatisticsApiRequestCycle) UnmarshalJSON(b []byte) error {
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
