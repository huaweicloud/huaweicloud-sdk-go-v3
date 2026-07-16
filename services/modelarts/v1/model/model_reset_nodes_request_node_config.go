package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResetNodesRequestNodeConfig **参数解释**：节点配置参数。 **约束限制**：该参数待下线。超节点禁传此参数，非超节点亦不推荐传递此参数。重置配置参数会从节点所在的nodepool自动取值， 若需要重置时修改某参数， 可先修改nodepool相关参数， 再进行重置。
type ResetNodesRequestNodeConfig struct {

	// **参数解释**：节点的镜像名称，如果不设置则取name字段的值 **约束限制**：不涉及。
	Os *string `json:"os,omitempty"`

	// **参数解释**：节点的镜像名称，如果os字段不设置才取此字段的值。 **约束限制**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：节点的镜像ID。 **约束限制**：不涉及。
	ImageId *string `json:"imageId,omitempty"`

	// **参数解释**：节点的镜像类型。 **约束限制**：不涉及。
	ImageType *string `json:"imageType,omitempty"`

	// **参数解释**：节点的容器运行时。 **约束限制**：不涉及。 **取值范围**：只能是[docker, containerd]其中一个。
	Runtime *ResetNodesRequestNodeConfigRuntime `json:"runtime,omitempty"`
}

func (o ResetNodesRequestNodeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodesRequestNodeConfig struct{}"
	}

	return strings.Join([]string{"ResetNodesRequestNodeConfig", string(data)}, " ")
}

type ResetNodesRequestNodeConfigRuntime struct {
	value string
}

type ResetNodesRequestNodeConfigRuntimeEnum struct {
	DOCKER     ResetNodesRequestNodeConfigRuntime
	CONTAINERD ResetNodesRequestNodeConfigRuntime
}

func GetResetNodesRequestNodeConfigRuntimeEnum() ResetNodesRequestNodeConfigRuntimeEnum {
	return ResetNodesRequestNodeConfigRuntimeEnum{
		DOCKER: ResetNodesRequestNodeConfigRuntime{
			value: "docker",
		},
		CONTAINERD: ResetNodesRequestNodeConfigRuntime{
			value: "containerd",
		},
	}
}

func (c ResetNodesRequestNodeConfigRuntime) Value() string {
	return c.value
}

func (c ResetNodesRequestNodeConfigRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetNodesRequestNodeConfigRuntime) UnmarshalJSON(b []byte) error {
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
