package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorMetadata **参数解释**：资源规格的元信息。
type ResourceFlavorMetadata struct {

	// **参数解释**：资源规格的ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	Labels *ResourceFlavorLabel `json:"labels"`

	Annotations *ResourceFlavorAnnotation `json:"annotations,omitempty"`
}

func (o ResourceFlavorMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorMetadata struct{}"
	}

	return strings.Join([]string{"ResourceFlavorMetadata", string(data)}, " ")
}
