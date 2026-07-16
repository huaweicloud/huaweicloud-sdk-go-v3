package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceFlavorsResponse Response Object
type ListResourceFlavorsResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - ResourceFlavorList：资源规格列表
	Kind *string `json:"kind,omitempty"`

	Metadata *ResourceFlavorListMetadata `json:"metadata,omitempty"`

	// 参数解释：资源规格列表。 取值范围：不涉及。
	Items          *[]ResourceFlavor `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListResourceFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceFlavorsResponse", string(data)}, " ")
}
