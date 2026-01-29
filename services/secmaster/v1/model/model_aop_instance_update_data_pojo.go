package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AopInstanceUpdateDataPojo 更新流程实例的请求对象
type AopInstanceUpdateDataPojo struct {

	// **参数解释**: 更新流程实例的指令 - ActionInstanceTerminateCommand 终止流程实例 - ActionInstanceRetryCommand 重试流程实例 - ActionInstanceDebugCommand 更新流程实例的调试结果  **约束限制**: 当command_type=ActionInstanceDebugCommand时task_id和inputdataobject是必填参数         **取值范围**: - ActionInstanceTerminateCommand - ActionInstanceRetryCommand - ActionInstanceDebugCommand  **默认值**:  不涉及
	CommandType AopInstanceUpdateDataPojoCommandType `json:"command_type"`

	// **参数解释**:         更新流程调试实例节点ID **约束限制**: 当command_type=ActionInstanceDebugCommand时参数为必填参数        **取值范围**: 不涉及 **默认值**:  不涉及
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 更新流程调试实例节点参数 **约束限制**: 当command_type=ActionInstanceDebugCommand时参数为必填参 **取值范围**: 不涉及 **默认值**:  不涉及
	InputDataobject *string `json:"input_dataobject,omitempty"`
}

func (o AopInstanceUpdateDataPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopInstanceUpdateDataPojo struct{}"
	}

	return strings.Join([]string{"AopInstanceUpdateDataPojo", string(data)}, " ")
}

type AopInstanceUpdateDataPojoCommandType struct {
	value string
}

type AopInstanceUpdateDataPojoCommandTypeEnum struct {
	ACTION_INSTANCE_TERMINATE_COMMAND AopInstanceUpdateDataPojoCommandType
	ACTION_INSTANCE_RETRY_COMMAND     AopInstanceUpdateDataPojoCommandType
	ACTION_INSTANCE_DEBUG_COMMAND     AopInstanceUpdateDataPojoCommandType
}

func GetAopInstanceUpdateDataPojoCommandTypeEnum() AopInstanceUpdateDataPojoCommandTypeEnum {
	return AopInstanceUpdateDataPojoCommandTypeEnum{
		ACTION_INSTANCE_TERMINATE_COMMAND: AopInstanceUpdateDataPojoCommandType{
			value: "ActionInstanceTerminateCommand",
		},
		ACTION_INSTANCE_RETRY_COMMAND: AopInstanceUpdateDataPojoCommandType{
			value: "ActionInstanceRetryCommand",
		},
		ACTION_INSTANCE_DEBUG_COMMAND: AopInstanceUpdateDataPojoCommandType{
			value: "ActionInstanceDebugCommand",
		},
	}
}

func (c AopInstanceUpdateDataPojoCommandType) Value() string {
	return c.value
}

func (c AopInstanceUpdateDataPojoCommandType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AopInstanceUpdateDataPojoCommandType) UnmarshalJSON(b []byte) error {
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
