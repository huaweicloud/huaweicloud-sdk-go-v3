package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 失败原因
type JobValidationResult struct {

	// 错误描述
	Message *string `json:"message,omitempty"`

	// ERROR,WARNING
	Status *JobValidationResultStatus `json:"status,omitempty"`
}

func (o JobValidationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobValidationResult struct{}"
	}

	return strings.Join([]string{"JobValidationResult", string(data)}, " ")
}

type JobValidationResultStatus struct {
	value string
}

type JobValidationResultStatusEnum struct {
	ERROR   JobValidationResultStatus
	WARNING JobValidationResultStatus
}

func GetJobValidationResultStatusEnum() JobValidationResultStatusEnum {
	return JobValidationResultStatusEnum{
		ERROR: JobValidationResultStatus{
			value: "ERROR",
		},
		WARNING: JobValidationResultStatus{
			value: "WARNING",
		},
	}
}

func (c JobValidationResultStatus) Value() string {
	return c.value
}

func (c JobValidationResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobValidationResultStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
