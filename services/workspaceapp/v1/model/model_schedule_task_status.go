package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScheduleTaskStatus job状态： * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 完成 * `FAILED` - 失败
type ScheduleTaskStatus struct {
	value string
}

type ScheduleTaskStatusEnum struct {
	WAITING ScheduleTaskStatus
	RUNNING ScheduleTaskStatus
	SUCCESS ScheduleTaskStatus
	FAILED  ScheduleTaskStatus
}

func GetScheduleTaskStatusEnum() ScheduleTaskStatusEnum {
	return ScheduleTaskStatusEnum{
		WAITING: ScheduleTaskStatus{
			value: "WAITING",
		},
		RUNNING: ScheduleTaskStatus{
			value: "RUNNING",
		},
		SUCCESS: ScheduleTaskStatus{
			value: "SUCCESS",
		},
		FAILED: ScheduleTaskStatus{
			value: "FAILED",
		},
	}
}

func (c ScheduleTaskStatus) Value() string {
	return c.value
}

func (c ScheduleTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleTaskStatus) UnmarshalJSON(b []byte) error {
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
