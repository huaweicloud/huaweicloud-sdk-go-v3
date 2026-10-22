package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeRayActorOptions **参数解释**：Actor 资源选项（可选，整体非必填）。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type ServeRayActorOptions struct {

	// **参数解释**：CPU 核数。 **约束限制**：不涉及。 **取值范围**：不小于0。 **默认取值**：不涉及。
	NumCpus float32 `json:"num_cpus,omitempty"`

	// **参数解释**：GPU 卡数。 **约束限制**：不涉及。 **取值范围**：不小于0。 **默认取值**：不涉及。
	NumGpus float32 `json:"num_gpus,omitempty"`

	// **参数解释**：自定义资源映射。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Resources map[string]float32 `json:"resources,omitempty"`
}

func (o ServeRayActorOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeRayActorOptions struct{}"
	}

	return strings.Join([]string{"ServeRayActorOptions", string(data)}, " ")
}
