package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AopInstanceEventData 操作流程实例的请求信息
type AopInstanceEventData struct {

	// **参数解释**: 操作流程实例的指令 - ActionInstanceRunCommand 运行流程实例  **约束限制**: 不涉及         **取值范围**: - ActionInstanceRunCommand  **默认值**:  不涉及
	CommandType AopInstanceEventDataCommandType `json:"command_type"`

	// **参数解释**: action的类型 - workflow 流程  **约束限制**: 不涉及         **取值范围**: - workflow  **默认值**:  不涉及
	ActionType AopInstanceEventDataActionType `json:"action_type"`

	// **参数解释**: 流程的ID **约束限制**: 不涉及
	ActionId *string `json:"action_id,omitempty"`

	// **参数解释**: 流程实例的ID **约束限制**: 不涉及
	ActionInstanceId *string `json:"action_instance_id,omitempty"`

	PlaybookContext *PlaybookcontextRef `json:"playbook_context,omitempty"`
}

func (o AopInstanceEventData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopInstanceEventData struct{}"
	}

	return strings.Join([]string{"AopInstanceEventData", string(data)}, " ")
}

type AopInstanceEventDataCommandType struct {
	value string
}

type AopInstanceEventDataCommandTypeEnum struct {
	ACTION_INSTANCE_RUN_COMMAND AopInstanceEventDataCommandType
}

func GetAopInstanceEventDataCommandTypeEnum() AopInstanceEventDataCommandTypeEnum {
	return AopInstanceEventDataCommandTypeEnum{
		ACTION_INSTANCE_RUN_COMMAND: AopInstanceEventDataCommandType{
			value: "ActionInstanceRunCommand",
		},
	}
}

func (c AopInstanceEventDataCommandType) Value() string {
	return c.value
}

func (c AopInstanceEventDataCommandType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AopInstanceEventDataCommandType) UnmarshalJSON(b []byte) error {
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

type AopInstanceEventDataActionType struct {
	value string
}

type AopInstanceEventDataActionTypeEnum struct {
	WORKFLOW AopInstanceEventDataActionType
}

func GetAopInstanceEventDataActionTypeEnum() AopInstanceEventDataActionTypeEnum {
	return AopInstanceEventDataActionTypeEnum{
		WORKFLOW: AopInstanceEventDataActionType{
			value: "workflow",
		},
	}
}

func (c AopInstanceEventDataActionType) Value() string {
	return c.value
}

func (c AopInstanceEventDataActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AopInstanceEventDataActionType) UnmarshalJSON(b []byte) error {
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
