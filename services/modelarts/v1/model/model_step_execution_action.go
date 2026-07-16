package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StepExecutionAction struct {

	// 操作名称，枚举如下:  - retry 重试  - stop 停止  - continue 继续
	ActionName StepExecutionActionActionName `json:"action_name"`

	DataRequirements *[]DataRequirement `json:"data_requirements,omitempty"`

	Parameters *[]WorkflowParameter `json:"parameters,omitempty"`
}

func (o StepExecutionAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepExecutionAction struct{}"
	}

	return strings.Join([]string{"StepExecutionAction", string(data)}, " ")
}

type StepExecutionActionActionName struct {
	value string
}

type StepExecutionActionActionNameEnum struct {
	RETRY    StepExecutionActionActionName
	STOP     StepExecutionActionActionName
	CONTINUE StepExecutionActionActionName
}

func GetStepExecutionActionActionNameEnum() StepExecutionActionActionNameEnum {
	return StepExecutionActionActionNameEnum{
		RETRY: StepExecutionActionActionName{
			value: "retry：重试",
		},
		STOP: StepExecutionActionActionName{
			value: "stop：停止",
		},
		CONTINUE: StepExecutionActionActionName{
			value: "continue：停止",
		},
	}
}

func (c StepExecutionActionActionName) Value() string {
	return c.value
}

func (c StepExecutionActionActionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepExecutionActionActionName) UnmarshalJSON(b []byte) error {
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
