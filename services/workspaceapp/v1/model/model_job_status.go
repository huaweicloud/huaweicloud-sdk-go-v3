package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobStatus job状态 * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 完成 * `FAILED` - 失败
type JobStatus struct {
	value string
}

type JobStatusEnum struct {
	WAITING JobStatus
	RUNNING JobStatus
	SUCCESS JobStatus
	FAILED  JobStatus
}

func GetJobStatusEnum() JobStatusEnum {
	return JobStatusEnum{
		WAITING: JobStatus{
			value: "WAITING",
		},
		RUNNING: JobStatus{
			value: "RUNNING",
		},
		SUCCESS: JobStatus{
			value: "SUCCESS",
		},
		FAILED: JobStatus{
			value: "FAILED",
		},
	}
}

func (c JobStatus) Value() string {
	return c.value
}

func (c JobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobStatus) UnmarshalJSON(b []byte) error {
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
