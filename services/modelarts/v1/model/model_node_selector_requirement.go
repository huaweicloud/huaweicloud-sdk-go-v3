package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeSelectorRequirement struct {

	// **参数解释**：选择器应用的标签键。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：表示键与一组值之间的关系。  **约束限制**：有效的运算符包括 In、NotIn、Exists、DoesNotExist、Gt 和 Lt。 **取值范围**： - In：表示键的值必须在给定的值列表中。例如，如果键是 color，值列表是 [\"red\", \"blue\"]，那么 color In [\"red\", \"blue\"] 表示 color 的值必须是 red 或 blue。 - NotIn：表示键的值不能在给定的值列表中。例如，color NotIn [\"red\", \"blue\"] 表示 color 的值不能是 red 或 blue。 - Exists：表示键必须存在，但对其值没有特定要求。例如，color Exists 表示必须存在 color 这个键，无论其值是什么。 - DoesNotExist：表示键不能存在。例如，color DoesNotExist 表示不能存在 color 这个键。 - Gt：表示键的值必须大于给定的值。例如，如果键是 age，age Gt 18 表示 age 的值必须大于 18。 - Lt：表示键的值必须小于给定的值。例如，age Lt 18 表示 age 的值必须小于 18。 **默认取值**：不涉及。
	Operator NodeSelectorRequirementOperator `json:"operator"`

	// **参数解释**：一个字符串值数组。 **约束限制**：如果操作符是“In”或“NotIn”，则该值数组不能为空。如果操作符是“Exists”或“DoesNotExist”，则该值数组必须为空。如果操作符是“Gt”或“Lt”，则该值数组必须包含一个元素，该元素将被解释为整数。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Values *[]string `json:"values,omitempty"`
}

func (o NodeSelectorRequirement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSelectorRequirement struct{}"
	}

	return strings.Join([]string{"NodeSelectorRequirement", string(data)}, " ")
}

type NodeSelectorRequirementOperator struct {
	value string
}

type NodeSelectorRequirementOperatorEnum struct {
	IN             NodeSelectorRequirementOperator
	NOT_IN         NodeSelectorRequirementOperator
	EXISTS         NodeSelectorRequirementOperator
	DOES_NOT_EXIST NodeSelectorRequirementOperator
	GT             NodeSelectorRequirementOperator
	LT             NodeSelectorRequirementOperator
}

func GetNodeSelectorRequirementOperatorEnum() NodeSelectorRequirementOperatorEnum {
	return NodeSelectorRequirementOperatorEnum{
		IN: NodeSelectorRequirementOperator{
			value: "In",
		},
		NOT_IN: NodeSelectorRequirementOperator{
			value: "NotIn",
		},
		EXISTS: NodeSelectorRequirementOperator{
			value: "Exists",
		},
		DOES_NOT_EXIST: NodeSelectorRequirementOperator{
			value: "DoesNotExist",
		},
		GT: NodeSelectorRequirementOperator{
			value: "Gt",
		},
		LT: NodeSelectorRequirementOperator{
			value: "Lt",
		},
	}
}

func (c NodeSelectorRequirementOperator) Value() string {
	return c.value
}

func (c NodeSelectorRequirementOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeSelectorRequirementOperator) UnmarshalJSON(b []byte) error {
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
