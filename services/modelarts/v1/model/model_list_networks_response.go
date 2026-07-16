package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworksResponse Response Object
type ListNetworksResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - NetworkList：网络列表。
	Kind *string `json:"kind,omitempty"`

	Metadata *NetworkListMetadata `json:"metadata,omitempty"`

	// **参数解释**：网络资源列表。
	Items          *[]Network `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListNetworksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworksResponse struct{}"
	}

	return strings.Join([]string{"ListNetworksResponse", string(data)}, " ")
}
