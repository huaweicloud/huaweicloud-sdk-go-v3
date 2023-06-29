package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AsyncTaskStatus 异步任务的状态：等待中、运行中、已完成、失败
type AsyncTaskStatus struct {
	value string
}

type AsyncTaskStatusEnum struct {
	WAITING  AsyncTaskStatus
	RUNNING  AsyncTaskStatus
	FINISHED AsyncTaskStatus
	FAILED   AsyncTaskStatus
}

func GetAsyncTaskStatusEnum() AsyncTaskStatusEnum {
	return AsyncTaskStatusEnum{
		WAITING: AsyncTaskStatus{
			value: "waiting",
		},
		RUNNING: AsyncTaskStatus{
			value: "running",
		},
		FINISHED: AsyncTaskStatus{
			value: "finished",
		},
		FAILED: AsyncTaskStatus{
			value: "failed",
		},
	}
}

func (c AsyncTaskStatus) Value() string {
	return c.value
}

func (c AsyncTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AsyncTaskStatus) UnmarshalJSON(b []byte) error {
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
