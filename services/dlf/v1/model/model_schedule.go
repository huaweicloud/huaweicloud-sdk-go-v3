package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Schedule struct {
	ScheType *ScheduleScheType `json:"scheType,omitempty"`

	Cron *Cron `json:"cron,omitempty"`

	Event *Event `json:"event,omitempty"`
}

func (o Schedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Schedule struct{}"
	}

	return strings.Join([]string{"Schedule", string(data)}, " ")
}

type ScheduleScheType struct {
	value string
}

type ScheduleScheTypeEnum struct {
	EXECUTE_ONCE ScheduleScheType
	CRON         ScheduleScheType
	EVENT        ScheduleScheType
}

func GetScheduleScheTypeEnum() ScheduleScheTypeEnum {
	return ScheduleScheTypeEnum{
		EXECUTE_ONCE: ScheduleScheType{
			value: "EXECUTE_ONCE",
		},
		CRON: ScheduleScheType{
			value: "CRON",
		},
		EVENT: ScheduleScheType{
			value: "EVENT",
		},
	}
}

func (c ScheduleScheType) Value() string {
	return c.value
}

func (c ScheduleScheType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleScheType) UnmarshalJSON(b []byte) error {
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
