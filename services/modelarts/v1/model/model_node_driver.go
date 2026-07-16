package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeDriver 节点驱动。
type NodeDriver struct {

	// **参数解释**：节点上驱动的版本号。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**：节点驱动升级策略。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UpdateStrategy *string `json:"updateStrategy,omitempty"`
}

func (o NodeDriver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDriver struct{}"
	}

	return strings.Join([]string{"NodeDriver", string(data)}, " ")
}
