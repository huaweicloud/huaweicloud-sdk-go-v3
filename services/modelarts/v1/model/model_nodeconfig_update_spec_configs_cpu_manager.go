package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeconfigUpdateSpecConfigsCpuManager **参数解释**：CPU管理策略配置，用于设置节点CPU的绑核管理方式。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigUpdateSpecConfigsCpuManager struct {

	// **参数解释**：CPU管理策略模式。 **约束限制**：不涉及。 **取值范围**：可选值如下： - static：静态绑核。 - none：不绑定。 **默认取值**：不涉及。
	Mode *NodeconfigUpdateSpecConfigsCpuManagerMode `json:"mode,omitempty"`
}

func (o NodeconfigUpdateSpecConfigsCpuManager) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdateSpecConfigsCpuManager struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdateSpecConfigsCpuManager", string(data)}, " ")
}

type NodeconfigUpdateSpecConfigsCpuManagerMode struct {
	value string
}

type NodeconfigUpdateSpecConfigsCpuManagerModeEnum struct {
	STATIC NodeconfigUpdateSpecConfigsCpuManagerMode
	NONE   NodeconfigUpdateSpecConfigsCpuManagerMode
}

func GetNodeconfigUpdateSpecConfigsCpuManagerModeEnum() NodeconfigUpdateSpecConfigsCpuManagerModeEnum {
	return NodeconfigUpdateSpecConfigsCpuManagerModeEnum{
		STATIC: NodeconfigUpdateSpecConfigsCpuManagerMode{
			value: "static",
		},
		NONE: NodeconfigUpdateSpecConfigsCpuManagerMode{
			value: "none",
		},
	}
}

func (c NodeconfigUpdateSpecConfigsCpuManagerMode) Value() string {
	return c.value
}

func (c NodeconfigUpdateSpecConfigsCpuManagerMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeconfigUpdateSpecConfigsCpuManagerMode) UnmarshalJSON(b []byte) error {
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
