package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeconfigUpdateSpecConfigsTransparentHugepage **参数解释**：透明大页（Transparent Hugepage，THP）配置，用于设置节点内存大页策略。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigUpdateSpecConfigsTransparentHugepage struct {

	// **参数解释**：透明大页模式。 **约束限制**：不涉及。 **取值范围**：可选值如下： - always：始终启用。 - madvise：建议启用。 - never：不启用。 **默认取值**：不涉及。
	Mode *NodeconfigUpdateSpecConfigsTransparentHugepageMode `json:"mode,omitempty"`
}

func (o NodeconfigUpdateSpecConfigsTransparentHugepage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdateSpecConfigsTransparentHugepage struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdateSpecConfigsTransparentHugepage", string(data)}, " ")
}

type NodeconfigUpdateSpecConfigsTransparentHugepageMode struct {
	value string
}

type NodeconfigUpdateSpecConfigsTransparentHugepageModeEnum struct {
	ALWAYS  NodeconfigUpdateSpecConfigsTransparentHugepageMode
	MADVISE NodeconfigUpdateSpecConfigsTransparentHugepageMode
	NEVER   NodeconfigUpdateSpecConfigsTransparentHugepageMode
}

func GetNodeconfigUpdateSpecConfigsTransparentHugepageModeEnum() NodeconfigUpdateSpecConfigsTransparentHugepageModeEnum {
	return NodeconfigUpdateSpecConfigsTransparentHugepageModeEnum{
		ALWAYS: NodeconfigUpdateSpecConfigsTransparentHugepageMode{
			value: "always",
		},
		MADVISE: NodeconfigUpdateSpecConfigsTransparentHugepageMode{
			value: "madvise",
		},
		NEVER: NodeconfigUpdateSpecConfigsTransparentHugepageMode{
			value: "never",
		},
	}
}

func (c NodeconfigUpdateSpecConfigsTransparentHugepageMode) Value() string {
	return c.value
}

func (c NodeconfigUpdateSpecConfigsTransparentHugepageMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeconfigUpdateSpecConfigsTransparentHugepageMode) UnmarshalJSON(b []byte) error {
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
