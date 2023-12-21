package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobDetailStatus job详情的状态 * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 成功 * `FAILED` - 失败 * `ABNORMAL` - 异常 * `ROLLBACK` - 回滚中 * `ABORTING` - 终止中
type JobDetailStatus struct {
	value string
}

type JobDetailStatusEnum struct {
	WAITING  JobDetailStatus
	RUNNING  JobDetailStatus
	SUCCESS  JobDetailStatus
	FAILED   JobDetailStatus
	ABNORMAL JobDetailStatus
	ROLLBACK JobDetailStatus
	ABORTING JobDetailStatus
}

func GetJobDetailStatusEnum() JobDetailStatusEnum {
	return JobDetailStatusEnum{
		WAITING: JobDetailStatus{
			value: "WAITING",
		},
		RUNNING: JobDetailStatus{
			value: "RUNNING",
		},
		SUCCESS: JobDetailStatus{
			value: "SUCCESS",
		},
		FAILED: JobDetailStatus{
			value: "FAILED",
		},
		ABNORMAL: JobDetailStatus{
			value: "ABNORMAL",
		},
		ROLLBACK: JobDetailStatus{
			value: "ROLLBACK",
		},
		ABORTING: JobDetailStatus{
			value: "ABORTING",
		},
	}
}

func (c JobDetailStatus) Value() string {
	return c.value
}

func (c JobDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobDetailStatus) UnmarshalJSON(b []byte) error {
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
