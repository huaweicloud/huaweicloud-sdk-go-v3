package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StepDetail 任务详情
type StepDetail struct {

	// 任务id
	StepId *string `json:"step_id,omitempty"`

	// 任务名
	StepName *string `json:"step_name,omitempty"`

	// 任务状态
	StepStatus *StepDetailStepStatus `json:"step_status,omitempty"`

	// 任务启动时间，格式为2020-06-17T07:38:42.503Z
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间，格式为2020-06-17T07:38:42.503Z
	EndTime *string `json:"end_time,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 子任务详情列表
	SubStepDetails *[]SubStepDetail `json:"sub_step_details,omitempty"`
}

func (o StepDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepDetail struct{}"
	}

	return strings.Join([]string{"StepDetail", string(data)}, " ")
}

type StepDetailStepStatus struct {
	value string
}

type StepDetailStepStatusEnum struct {
	FINISH    StepDetailStepStatus
	FAILED    StepDetailStepStatus
	EXECUTING StepDetailStepStatus
	WAITING   StepDetailStepStatus
}

func GetStepDetailStepStatusEnum() StepDetailStepStatusEnum {
	return StepDetailStepStatusEnum{
		FINISH: StepDetailStepStatus{
			value: "FINISH",
		},
		FAILED: StepDetailStepStatus{
			value: "FAILED",
		},
		EXECUTING: StepDetailStepStatus{
			value: "EXECUTING",
		},
		WAITING: StepDetailStepStatus{
			value: "WAITING",
		},
	}
}

func (c StepDetailStepStatus) Value() string {
	return c.value
}

func (c StepDetailStepStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepDetailStepStatus) UnmarshalJSON(b []byte) error {
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
