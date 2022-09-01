package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 函数执行节点信息
type OperationState struct {

	// Action执行模式，支持串行，并行两种模式，默认串行
	ActionMode *OperationStateActionMode `json:"action_mode,omitempty" xml:"action_mode"`

	// 节点中要执行的操作列表
	Actions []Action `json:"actions" xml:"actions"`

	// 错误处理策略
	OnErrors *[]OnError `json:"on_errors,omitempty" xml:"on_errors"`

	// 节点ID，需要在当前函数流中唯一
	Id string `json:"id" xml:"id"`

	// 节点名称
	Name string `json:"name" xml:"name"`

	// 节点类型
	Type OperationStateType `json:"type" xml:"type"`

	// 是否是结束节点
	End bool `json:"end" xml:"end"`

	// 下一步骤节点ID
	Transition string `json:"transition" xml:"transition"`

	StateDataFilter *StateDataFilter `json:"state_data_filter,omitempty" xml:"state_data_filter"`

	// 时间等待节点等待时间（秒）,节点类型为Sleep时为必填，节点类型不为Sleep时无效
	Duration *int64 `json:"duration,omitempty" xml:"duration"`
}

func (o OperationState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationState struct{}"
	}

	return strings.Join([]string{"OperationState", string(data)}, " ")
}

type OperationStateActionMode struct {
	value string
}

type OperationStateActionModeEnum struct {
	SEQUENTIAL OperationStateActionMode
	PARALLEL   OperationStateActionMode
}

func GetOperationStateActionModeEnum() OperationStateActionModeEnum {
	return OperationStateActionModeEnum{
		SEQUENTIAL: OperationStateActionMode{
			value: "sequential",
		},
		PARALLEL: OperationStateActionMode{
			value: "parallel",
		},
	}
}

func (c OperationStateActionMode) Value() string {
	return c.value
}

func (c OperationStateActionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationStateActionMode) UnmarshalJSON(b []byte) error {
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

type OperationStateType struct {
	value string
}

type OperationStateTypeEnum struct {
	OPERATION OperationStateType
	SLEEP     OperationStateType
	END       OperationStateType
}

func GetOperationStateTypeEnum() OperationStateTypeEnum {
	return OperationStateTypeEnum{
		OPERATION: OperationStateType{
			value: "Operation",
		},
		SLEEP: OperationStateType{
			value: "Sleep",
		},
		END: OperationStateType{
			value: "End",
		},
	}
}

func (c OperationStateType) Value() string {
	return c.value
}

func (c OperationStateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationStateType) UnmarshalJSON(b []byte) error {
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
