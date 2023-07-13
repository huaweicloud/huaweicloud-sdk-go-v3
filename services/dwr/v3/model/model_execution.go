package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Execution struct {

	// 工作流URN
	GraphUrn *string `json:"graph_urn,omitempty"`

	// 工作流实例URN
	ExecutionUrn *string `json:"execution_urn,omitempty"`

	// 工作流实例启动时间
	StartedAt *string `json:"started_at,omitempty"`

	// 工作流执行方式。APICALL代表为通过API方式触发。
	ExecutionType *ExecutionExecutionType `json:"execution_type,omitempty"`

	// 工作流停止时间。
	StoppedAt *string `json:"stopped_at,omitempty"`

	// 工作流运行状态。success,fail,running,timeout,cancel
	ExecutionState *ExecutionExecutionState `json:"execution_state,omitempty"`

	// 工作流名称。
	ExecutionName *string `json:"execution_name,omitempty"`
}

func (o Execution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Execution struct{}"
	}

	return strings.Join([]string{"Execution", string(data)}, " ")
}

type ExecutionExecutionType struct {
	value string
}

type ExecutionExecutionTypeEnum struct {
	APICALL ExecutionExecutionType
}

func GetExecutionExecutionTypeEnum() ExecutionExecutionTypeEnum {
	return ExecutionExecutionTypeEnum{
		APICALL: ExecutionExecutionType{
			value: "APICALL",
		},
	}
}

func (c ExecutionExecutionType) Value() string {
	return c.value
}

func (c ExecutionExecutionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionExecutionType) UnmarshalJSON(b []byte) error {
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

type ExecutionExecutionState struct {
	value string
}

type ExecutionExecutionStateEnum struct {
	SUCCESS ExecutionExecutionState
	FAIL    ExecutionExecutionState
	RUNNING ExecutionExecutionState
	TIMEOUT ExecutionExecutionState
	CANCEL  ExecutionExecutionState
}

func GetExecutionExecutionStateEnum() ExecutionExecutionStateEnum {
	return ExecutionExecutionStateEnum{
		SUCCESS: ExecutionExecutionState{
			value: "success",
		},
		FAIL: ExecutionExecutionState{
			value: "fail",
		},
		RUNNING: ExecutionExecutionState{
			value: "running",
		},
		TIMEOUT: ExecutionExecutionState{
			value: "timeout",
		},
		CANCEL: ExecutionExecutionState{
			value: "cancel",
		},
	}
}

func (c ExecutionExecutionState) Value() string {
	return c.value
}

func (c ExecutionExecutionState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionExecutionState) UnmarshalJSON(b []byte) error {
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
