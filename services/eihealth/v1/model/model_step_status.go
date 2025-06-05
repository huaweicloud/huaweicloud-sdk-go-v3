package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StepStatus struct {
	value string
}

type StepStatusEnum struct {
	RUNNING  StepStatus
	FAILED   StepStatus
	FINISHED StepStatus
}

func GetStepStatusEnum() StepStatusEnum {
	return StepStatusEnum{
		RUNNING: StepStatus{
			value: "RUNNING",
		},
		FAILED: StepStatus{
			value: "FAILED",
		},
		FINISHED: StepStatus{
			value: "FINISHED",
		},
	}
}

func (c StepStatus) Value() string {
	return c.value
}

func (c StepStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepStatus) UnmarshalJSON(b []byte) error {
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
