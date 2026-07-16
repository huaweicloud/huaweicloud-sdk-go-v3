package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNetworkResponse Response Object
type DeleteNetworkResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Network：网络。
	Kind *string `json:"kind,omitempty"`

	Metadata *NetworkMetadata `json:"metadata,omitempty"`

	Spec *NetworkSpec `json:"spec,omitempty"`

	Status         *NetworkStatus `json:"status,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DeleteNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNetworkResponse struct{}"
	}

	return strings.Join([]string{"DeleteNetworkResponse", string(data)}, " ")
}
