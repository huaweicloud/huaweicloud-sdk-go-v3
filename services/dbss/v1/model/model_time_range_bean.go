package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TimeRangeBean struct {

	// 开始时间，必须和end_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	EndTime *string `json:"end_time,omitempty"`

	// 结束时间，必须和start_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	StartTime *string `json:"start_time,omitempty"`

	// 请求查询的时间段，和start_time，end_time不能同时使用，同时传该参数优先级更高。 - HALF_HOUR - HOUR - THREE_HOUR - TWELVE_HOUR - DAY - WEEK - MONTH
	TimeRange *TimeRangeBeanTimeRange `json:"time_range,omitempty"`
}

func (o TimeRangeBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRangeBean struct{}"
	}

	return strings.Join([]string{"TimeRangeBean", string(data)}, " ")
}

type TimeRangeBeanTimeRange struct {
	value string
}

type TimeRangeBeanTimeRangeEnum struct {
	HALF_HOUR   TimeRangeBeanTimeRange
	HOUR        TimeRangeBeanTimeRange
	THREE_HOUR  TimeRangeBeanTimeRange
	TWELVE_HOUR TimeRangeBeanTimeRange
	DAY         TimeRangeBeanTimeRange
	WEEK        TimeRangeBeanTimeRange
	MONTH       TimeRangeBeanTimeRange
}

func GetTimeRangeBeanTimeRangeEnum() TimeRangeBeanTimeRangeEnum {
	return TimeRangeBeanTimeRangeEnum{
		HALF_HOUR: TimeRangeBeanTimeRange{
			value: "HALF_HOUR",
		},
		HOUR: TimeRangeBeanTimeRange{
			value: " HOUR",
		},
		THREE_HOUR: TimeRangeBeanTimeRange{
			value: " THREE_HOUR",
		},
		TWELVE_HOUR: TimeRangeBeanTimeRange{
			value: " TWELVE_HOUR",
		},
		DAY: TimeRangeBeanTimeRange{
			value: " DAY",
		},
		WEEK: TimeRangeBeanTimeRange{
			value: " WEEK",
		},
		MONTH: TimeRangeBeanTimeRange{
			value: " MONTH",
		},
	}
}

func (c TimeRangeBeanTimeRange) Value() string {
	return c.value
}

func (c TimeRangeBeanTimeRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TimeRangeBeanTimeRange) UnmarshalJSON(b []byte) error {
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
