package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorSpecGpu **参数解释**：资源规格实例的GPU资源信息。
type ResourceFlavorSpecGpu struct {

	// **参数解释**：资源规格实例的GPU卡类型。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：资源规格实例的GPU卡数量。 **取值范围**：不涉及。
	Size *string `json:"size,omitempty"`
}

func (o ResourceFlavorSpecGpu) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorSpecGpu struct{}"
	}

	return strings.Join([]string{"ResourceFlavorSpecGpu", string(data)}, " ")
}
