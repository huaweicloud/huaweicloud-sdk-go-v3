package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WorkflowParameterResp 参数。
type WorkflowParameterResp struct {

	// **参数解释**：Workflow工作流配置参数的名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：参数的类型。 **取值范围**：枚举值如下: - str：字符串 - int：整型 - bool：布尔类型 - float：浮点型
	Type *WorkflowParameterRespType `json:"type,omitempty"`

	// **参数解释**：Workflow工作流配置参数的描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：Workflow工作流配置参数的样例。
	Example *interface{} `json:"example,omitempty"`

	// **参数解释**：是否为延迟输入的参数。 **取值范围**： - true：是 - false：否
	Delay *bool `json:"delay,omitempty"`

	// **参数解释**：配置参数的默认值。
	Default *interface{} `json:"default,omitempty"`

	// **参数解释**：参数值。
	Value *interface{} `json:"value,omitempty"`

	// **参数解释**：Workflow工作流配置参数的枚举项。
	Enum *[]interface{} `json:"enum,omitempty"`

	// **参数解释**：使用这个参数的工作流节点。
	UsedSteps *[]string `json:"used_steps,omitempty"`

	// **参数解释**：数据格式。 **取值范围**：不涉及。
	Format *string `json:"format,omitempty"`

	// **参数解释**：限制条件。
	Constraint map[string]interface{} `json:"constraint,omitempty"`
}

func (o WorkflowParameterResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowParameterResp struct{}"
	}

	return strings.Join([]string{"WorkflowParameterResp", string(data)}, " ")
}

type WorkflowParameterRespType struct {
	value string
}

type WorkflowParameterRespTypeEnum struct {
	STR   WorkflowParameterRespType
	INT   WorkflowParameterRespType
	BOOL  WorkflowParameterRespType
	FLOAT WorkflowParameterRespType
}

func GetWorkflowParameterRespTypeEnum() WorkflowParameterRespTypeEnum {
	return WorkflowParameterRespTypeEnum{
		STR: WorkflowParameterRespType{
			value: "str：字符串类型",
		},
		INT: WorkflowParameterRespType{
			value: "int：整型",
		},
		BOOL: WorkflowParameterRespType{
			value: "bool：布尔类型",
		},
		FLOAT: WorkflowParameterRespType{
			value: "float：浮点型",
		},
	}
}

func (c WorkflowParameterRespType) Value() string {
	return c.value
}

func (c WorkflowParameterRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkflowParameterRespType) UnmarshalJSON(b []byte) error {
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
