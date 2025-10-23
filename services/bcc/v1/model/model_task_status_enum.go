package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskStatusEnum 任务状态枚举
type TaskStatusEnum struct {
	value string
}

type TaskStatusEnumEnum struct {
	SUCCESS TaskStatusEnum
	FAILED  TaskStatusEnum
	RUNNING TaskStatusEnum
	SKIPPED TaskStatusEnum
	TIMEOUT TaskStatusEnum
}

func GetTaskStatusEnumEnum() TaskStatusEnumEnum {
	return TaskStatusEnumEnum{
		SUCCESS: TaskStatusEnum{
			value: "success",
		},
		FAILED: TaskStatusEnum{
			value: "failed",
		},
		RUNNING: TaskStatusEnum{
			value: "running",
		},
		SKIPPED: TaskStatusEnum{
			value: "skipped",
		},
		TIMEOUT: TaskStatusEnum{
			value: "timeout",
		},
	}
}

func (c TaskStatusEnum) Value() string {
	return c.value
}

func (c TaskStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskStatusEnum) UnmarshalJSON(b []byte) error {
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
