package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WorkflowParameter 参数。
type WorkflowParameter struct {

	// Workflow工作流配置参数的名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	Name *string `json:"name,omitempty"`

	// 参数的类型，枚举值如下: - str：字符串 - int：整型 - bool：布尔类型 - float：浮点型
	Type *WorkflowParameterType `json:"type,omitempty"`

	// Workflow工作流配置参数的描述。
	Description *string `json:"description,omitempty"`

	// Workflow工作流配置参数的样例。
	Example *interface{} `json:"example,omitempty"`

	// 是否为延迟输入的参数，默认为否。
	Delay *bool `json:"delay,omitempty"`

	// 配置参数的默认值。
	Default *interface{} `json:"default,omitempty"`

	// 参数值。
	Value *interface{} `json:"value,omitempty"`

	// Workflow工作流配置参数的枚举项。
	Enum *[]interface{} `json:"enum,omitempty"`

	// 使用这个参数的工作流节点。
	UsedSteps *[]string `json:"used_steps,omitempty"`

	// 数据格式。
	Format *string `json:"format,omitempty"`

	// 限制条件。
	Constraint map[string]interface{} `json:"constraint,omitempty"`
}

func (o WorkflowParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowParameter struct{}"
	}

	return strings.Join([]string{"WorkflowParameter", string(data)}, " ")
}

type WorkflowParameterType struct {
	value string
}

type WorkflowParameterTypeEnum struct {
	STR   WorkflowParameterType
	INT   WorkflowParameterType
	BOOL  WorkflowParameterType
	FLOAT WorkflowParameterType
}

func GetWorkflowParameterTypeEnum() WorkflowParameterTypeEnum {
	return WorkflowParameterTypeEnum{
		STR: WorkflowParameterType{
			value: "str：字符串类型",
		},
		INT: WorkflowParameterType{
			value: "int：整型",
		},
		BOOL: WorkflowParameterType{
			value: "bool：布尔类型",
		},
		FLOAT: WorkflowParameterType{
			value: "float：浮点型",
		},
	}
}

func (c WorkflowParameterType) Value() string {
	return c.value
}

func (c WorkflowParameterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkflowParameterType) UnmarshalJSON(b []byte) error {
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
