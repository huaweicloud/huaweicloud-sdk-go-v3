package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobSubmitResult struct {

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业提交状态。 枚举值： - COMPLETE：作业提交完成。 - FAILED：作业提交失败。
	State *JobSubmitResultState `json:"state,omitempty"`
}

func (o JobSubmitResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSubmitResult struct{}"
	}

	return strings.Join([]string{"JobSubmitResult", string(data)}, " ")
}

type JobSubmitResultState struct {
	value string
}

type JobSubmitResultStateEnum struct {
	COMPLETE JobSubmitResultState
	FAILED   JobSubmitResultState
}

func GetJobSubmitResultStateEnum() JobSubmitResultStateEnum {
	return JobSubmitResultStateEnum{
		COMPLETE: JobSubmitResultState{
			value: "COMPLETE",
		},
		FAILED: JobSubmitResultState{
			value: "FAILED",
		},
	}
}

func (c JobSubmitResultState) Value() string {
	return c.value
}

func (c JobSubmitResultState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobSubmitResultState) UnmarshalJSON(b []byte) error {
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
