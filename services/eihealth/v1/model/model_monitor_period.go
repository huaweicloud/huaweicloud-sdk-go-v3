package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MonitorPeriod struct {
	value string
}

type MonitorPeriodEnum struct {
	REAL_TIME                 MonitorPeriod
	FIVE_MINUTES              MonitorPeriod
	FIFTEEN_TO_TWENTY_MINUTES MonitorPeriod
	ONE_HOUR                  MonitorPeriod
}

func GetMonitorPeriodEnum() MonitorPeriodEnum {
	return MonitorPeriodEnum{
		REAL_TIME: MonitorPeriod{
			value: "real_time",
		},
		FIVE_MINUTES: MonitorPeriod{
			value: "five_minutes",
		},
		FIFTEEN_TO_TWENTY_MINUTES: MonitorPeriod{
			value: "fifteen_to_twenty_minutes",
		},
		ONE_HOUR: MonitorPeriod{
			value: "one_hour",
		},
	}
}

func (c MonitorPeriod) Value() string {
	return c.value
}

func (c MonitorPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorPeriod) UnmarshalJSON(b []byte) error {
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
