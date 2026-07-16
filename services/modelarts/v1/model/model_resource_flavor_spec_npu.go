package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorSpecNpu **参数解释**：资源规格实例的NPU资源信息。
type ResourceFlavorSpecNpu struct {

	// **参数解释**：资源规格实例的NPU卡类型。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：资源规格实例的NPU卡数量。 **取值范围**：不涉及。
	Size *string `json:"size,omitempty"`
}

func (o ResourceFlavorSpecNpu) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorSpecNpu struct{}"
	}

	return strings.Join([]string{"ResourceFlavorSpecNpu", string(data)}, " ")
}
