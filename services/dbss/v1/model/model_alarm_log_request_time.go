package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmLogRequestTime 时间
type AlarmLogRequestTime struct {

	// 时间范围。和start_time，end_time不能同时使用，同时传该参数优先级更高。枚举值 HALF_HOUR, HOUR, THREE_HOUR, TWELVE_HOUR, DAY, WEEK, MONTH;
	TimeRange *AlarmLogRequestTimeTimeRange `json:"time_range,omitempty"`

	// 开始时间，必须和end_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，必须和start_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o AlarmLogRequestTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLogRequestTime struct{}"
	}

	return strings.Join([]string{"AlarmLogRequestTime", string(data)}, " ")
}

type AlarmLogRequestTimeTimeRange struct {
	value string
}

type AlarmLogRequestTimeTimeRangeEnum struct {
	HALF_HOUR   AlarmLogRequestTimeTimeRange
	HOUR        AlarmLogRequestTimeTimeRange
	THREE_HOUR  AlarmLogRequestTimeTimeRange
	TWELVE_HOUR AlarmLogRequestTimeTimeRange
	DAY         AlarmLogRequestTimeTimeRange
	WEEK        AlarmLogRequestTimeTimeRange
	MONTH       AlarmLogRequestTimeTimeRange
}

func GetAlarmLogRequestTimeTimeRangeEnum() AlarmLogRequestTimeTimeRangeEnum {
	return AlarmLogRequestTimeTimeRangeEnum{
		HALF_HOUR: AlarmLogRequestTimeTimeRange{
			value: "HALF_HOUR",
		},
		HOUR: AlarmLogRequestTimeTimeRange{
			value: " HOUR",
		},
		THREE_HOUR: AlarmLogRequestTimeTimeRange{
			value: " THREE_HOUR",
		},
		TWELVE_HOUR: AlarmLogRequestTimeTimeRange{
			value: " TWELVE_HOUR",
		},
		DAY: AlarmLogRequestTimeTimeRange{
			value: " DAY",
		},
		WEEK: AlarmLogRequestTimeTimeRange{
			value: " WEEK",
		},
		MONTH: AlarmLogRequestTimeTimeRange{
			value: " MONTH;",
		},
	}
}

func (c AlarmLogRequestTimeTimeRange) Value() string {
	return c.value
}

func (c AlarmLogRequestTimeTimeRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmLogRequestTimeTimeRange) UnmarshalJSON(b []byte) error {
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
