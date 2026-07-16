package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetNodesRequestRollingConfig **参数解释**：节点重置是滚动配置。 **约束限制**：不涉及。
type ResetNodesRequestRollingConfig struct {

	// **参数解释**：滚动策略。 **约束限制**：不涉及。 **取值范围**：可选值如下： - RollingByNumber：表示按节点数量设置最大同时重置节点数量，例如10，表示单次最多重置10个节点 - RollingByPercent：表示按百分比设置最大同时重置节点数量。例如10，表示单次最多重置10%的节点 **默认取值**：不涉及。
	Strategy string `json:"strategy"`

	// **参数解释**：滚动重置的节点数量或者节点比例, 当strategy为RollingByNumber时,表示允许同时重置的节点数量, 当strategy为RollingByPercent时,表示允许同时重置的最大节点比例。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MaxUnavailable int32 `json:"maxUnavailable"`
}

func (o ResetNodesRequestRollingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodesRequestRollingConfig struct{}"
	}

	return strings.Join([]string{"ResetNodesRequestRollingConfig", string(data)}, " ")
}
