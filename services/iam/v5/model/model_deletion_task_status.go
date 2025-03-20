package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeletionTaskStatus 删除任务状态。
type DeletionTaskStatus struct {
	value string
}

type DeletionTaskStatusEnum struct {
	SUCCEEDED   DeletionTaskStatus
	IN_PROGRESS DeletionTaskStatus
	FAILED      DeletionTaskStatus
	NOT_STARTED DeletionTaskStatus
}

func GetDeletionTaskStatusEnum() DeletionTaskStatusEnum {
	return DeletionTaskStatusEnum{
		SUCCEEDED: DeletionTaskStatus{
			value: "succeeded",
		},
		IN_PROGRESS: DeletionTaskStatus{
			value: "in_progress",
		},
		FAILED: DeletionTaskStatus{
			value: "failed",
		},
		NOT_STARTED: DeletionTaskStatus{
			value: "not_started",
		},
	}
}

func (c DeletionTaskStatus) Value() string {
	return c.value
}

func (c DeletionTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeletionTaskStatus) UnmarshalJSON(b []byte) error {
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
