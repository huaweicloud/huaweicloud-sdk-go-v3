package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeconfigUpdateSpecConfigsDropCache **参数解释**：节点缓存清理配置，用于定期清理节点内存缓存，释放内存空间。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigUpdateSpecConfigsDropCache struct {

	// **参数解释**：缓存清理开关。 **约束限制**：不涉及。 **取值范围**：可选值如下： - enable：启用。 - disable：禁用。 **默认取值**：不涉及。
	Mode *NodeconfigUpdateSpecConfigsDropCacheMode `json:"mode,omitempty"`
}

func (o NodeconfigUpdateSpecConfigsDropCache) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdateSpecConfigsDropCache struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdateSpecConfigsDropCache", string(data)}, " ")
}

type NodeconfigUpdateSpecConfigsDropCacheMode struct {
	value string
}

type NodeconfigUpdateSpecConfigsDropCacheModeEnum struct {
	ENABLE  NodeconfigUpdateSpecConfigsDropCacheMode
	DISABLE NodeconfigUpdateSpecConfigsDropCacheMode
}

func GetNodeconfigUpdateSpecConfigsDropCacheModeEnum() NodeconfigUpdateSpecConfigsDropCacheModeEnum {
	return NodeconfigUpdateSpecConfigsDropCacheModeEnum{
		ENABLE: NodeconfigUpdateSpecConfigsDropCacheMode{
			value: "enable",
		},
		DISABLE: NodeconfigUpdateSpecConfigsDropCacheMode{
			value: "disable",
		},
	}
}

func (c NodeconfigUpdateSpecConfigsDropCacheMode) Value() string {
	return c.value
}

func (c NodeconfigUpdateSpecConfigsDropCacheMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeconfigUpdateSpecConfigsDropCacheMode) UnmarshalJSON(b []byte) error {
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
