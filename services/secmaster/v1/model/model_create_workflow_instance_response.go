package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateWorkflowInstanceResponse Response Object
type CreateWorkflowInstanceResponse struct {

	// **参数解释**: 操作流程实例的指令 - ActionInstanceRunCommand 运行流程实例  **约束限制**: 不涉及         **取值范围**: - ActionInstanceRunCommand  **默认值**:  不涉及
	CommandType *CreateWorkflowInstanceResponseCommandType `json:"command_type,omitempty"`

	// **参数解释**: action的类型 - workflow 流程  **约束限制**: 不涉及         **取值范围**: - workflow  **默认值**:  不涉及
	ActionType *CreateWorkflowInstanceResponseActionType `json:"action_type,omitempty"`

	// **参数解释**: 流程的ID **约束限制**: 不涉及
	ActionId *string `json:"action_id,omitempty"`

	// **参数解释**: 流程实例的ID **约束限制**: 不涉及
	ActionInstanceId *string `json:"action_instance_id,omitempty"`

	PlaybookContext *PlaybookcontextRef `json:"playbook_context,omitempty"`
	HttpStatusCode  int                 `json:"-"`
}

func (o CreateWorkflowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowInstanceResponse", string(data)}, " ")
}

type CreateWorkflowInstanceResponseCommandType struct {
	value string
}

type CreateWorkflowInstanceResponseCommandTypeEnum struct {
	ACTION_INSTANCE_RUN_COMMAND CreateWorkflowInstanceResponseCommandType
}

func GetCreateWorkflowInstanceResponseCommandTypeEnum() CreateWorkflowInstanceResponseCommandTypeEnum {
	return CreateWorkflowInstanceResponseCommandTypeEnum{
		ACTION_INSTANCE_RUN_COMMAND: CreateWorkflowInstanceResponseCommandType{
			value: "ActionInstanceRunCommand",
		},
	}
}

func (c CreateWorkflowInstanceResponseCommandType) Value() string {
	return c.value
}

func (c CreateWorkflowInstanceResponseCommandType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWorkflowInstanceResponseCommandType) UnmarshalJSON(b []byte) error {
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

type CreateWorkflowInstanceResponseActionType struct {
	value string
}

type CreateWorkflowInstanceResponseActionTypeEnum struct {
	WORKFLOW CreateWorkflowInstanceResponseActionType
}

func GetCreateWorkflowInstanceResponseActionTypeEnum() CreateWorkflowInstanceResponseActionTypeEnum {
	return CreateWorkflowInstanceResponseActionTypeEnum{
		WORKFLOW: CreateWorkflowInstanceResponseActionType{
			value: "workflow",
		},
	}
}

func (c CreateWorkflowInstanceResponseActionType) Value() string {
	return c.value
}

func (c CreateWorkflowInstanceResponseActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWorkflowInstanceResponseActionType) UnmarshalJSON(b []byte) error {
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
