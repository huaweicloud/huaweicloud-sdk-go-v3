package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreferredSchedulingTerm struct {

	// **参数解释**：与匹配相应 nodeSelectorTerm 相关的权重 **约束限制**：不涉及。 **取值范围**：范围为 1-100。 **默认取值**：不涉及。
	Weight *int32 `json:"weight,omitempty"`

	Preference *NodeSelectorTerm `json:"preference,omitempty"`
}

func (o PreferredSchedulingTerm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreferredSchedulingTerm struct{}"
	}

	return strings.Join([]string{"PreferredSchedulingTerm", string(data)}, " ")
}
