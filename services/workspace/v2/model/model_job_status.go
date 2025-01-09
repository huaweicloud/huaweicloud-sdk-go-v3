package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobStatus 实例类型： * `INIT` - 初始化中 * `WAITING` - 等待安装结束 * `SUCCESS` - 成功 * `FAIL` - 失败
type JobStatus struct {
	value string
}

type JobStatusEnum struct {
	INIT    JobStatus
	WAITING JobStatus
	SUCCESS JobStatus
	FAIL    JobStatus
}

func GetJobStatusEnum() JobStatusEnum {
	return JobStatusEnum{
		INIT: JobStatus{
			value: "INIT",
		},
		WAITING: JobStatus{
			value: "WAITING",
		},
		SUCCESS: JobStatus{
			value: "SUCCESS",
		},
		FAIL: JobStatus{
			value: "FAIL",
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
