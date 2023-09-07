package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScheduledTaskPolicy struct {

	// 非必选，仅当recurrence_type不为空时生效，表示计划任务的生效开始时间，格式为yyyy-MM-dd'T'HH:mm'Z'，不填写时默认为任务创建成功的时间
	StartTime *string `json:"start_time,omitempty"`

	// 仅当recurrence_type不为空时生效且必选，表示计划任务的生效结束时间，格式为yyyy-MM-dd'T'HH:mm'Z'
	EndTime *string `json:"end_time,omitempty"`

	// 必选，执行时间，采用UTC时间，recurrence_type不填写或为空时，格式为HH:mm, recurrence_type不为空时，格式为yyyy-MM-dd'T'HH:mm'Z'
	LaunchTime string `json:"launch_time"`

	// 非必选，不填写时计划任务为定时执行， 填写时，为周期执行，且只能填写DAILY，WEEKLY，MONTHLY 中的一种，分别为按天，按周，按月周期执行
	RecurrenceType *ScheduledTaskPolicyRecurrenceType `json:"recurrence_type,omitempty"`

	// 仅当recurrence_type为WEEKLY，MONTHLY时必选，表示周期执行的具体日期，多个日期用,分割。recurrence_type为WEEKLY时，可填入1-7， recurrence_type为MONTHLY时，可填入1-31
	RecurrenceValue *string `json:"recurrence_value,omitempty"`
}

func (o ScheduledTaskPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskPolicy struct{}"
	}

	return strings.Join([]string{"ScheduledTaskPolicy", string(data)}, " ")
}

type ScheduledTaskPolicyRecurrenceType struct {
	value string
}

type ScheduledTaskPolicyRecurrenceTypeEnum struct {
	DAILY   ScheduledTaskPolicyRecurrenceType
	WEEKLY  ScheduledTaskPolicyRecurrenceType
	MONTHLY ScheduledTaskPolicyRecurrenceType
}

func GetScheduledTaskPolicyRecurrenceTypeEnum() ScheduledTaskPolicyRecurrenceTypeEnum {
	return ScheduledTaskPolicyRecurrenceTypeEnum{
		DAILY: ScheduledTaskPolicyRecurrenceType{
			value: "DAILY",
		},
		WEEKLY: ScheduledTaskPolicyRecurrenceType{
			value: "WEEKLY",
		},
		MONTHLY: ScheduledTaskPolicyRecurrenceType{
			value: "MONTHLY",
		},
	}
}

func (c ScheduledTaskPolicyRecurrenceType) Value() string {
	return c.value
}

func (c ScheduledTaskPolicyRecurrenceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledTaskPolicyRecurrenceType) UnmarshalJSON(b []byte) error {
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
