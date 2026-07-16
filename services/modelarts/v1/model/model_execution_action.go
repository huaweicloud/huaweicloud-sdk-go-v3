package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecutionAction 操作Workflow Execution的动作。
type ExecutionAction struct {

	// 操作名称，枚举如下: - stop 停止 - rerun 重跑
	ActionName ExecutionActionActionName `json:"action_name"`

	Policies *ExecutionActionPolicy `json:"policies,omitempty"`

	// 参数。
	Parameters *[]WorkflowParameter `json:"parameters,omitempty"`

	// 需要的数据。
	DataRequirements *[]DataRequirement `json:"data_requirements,omitempty"`
}

func (o ExecutionAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionAction struct{}"
	}

	return strings.Join([]string{"ExecutionAction", string(data)}, " ")
}

type ExecutionActionActionName struct {
	value string
}

type ExecutionActionActionNameEnum struct {
	STOP ExecutionActionActionName
}

func GetExecutionActionActionNameEnum() ExecutionActionActionNameEnum {
	return ExecutionActionActionNameEnum{
		STOP: ExecutionActionActionName{
			value: "stop",
		},
	}
}

func (c ExecutionActionActionName) Value() string {
	return c.value
}

func (c ExecutionActionActionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionActionActionName) UnmarshalJSON(b []byte) error {
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
