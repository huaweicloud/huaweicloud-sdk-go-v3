package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigSpecConfigsCpuManager **参数解释**：CPU管理策略配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigSpecConfigsCpuManager struct {

	// **参数解释**：CPU管理策略模式。 **约束限制**：不涉及。 **取值范围**：static、none。 **默认取值**：不涉及。
	Mode *string `json:"mode,omitempty"`
}

func (o NodeconfigSpecConfigsCpuManager) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigSpecConfigsCpuManager struct{}"
	}

	return strings.Join([]string{"NodeconfigSpecConfigsCpuManager", string(data)}, " ")
}
