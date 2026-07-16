package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolResponse Response Object
type ShowPoolResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v2：当前资源版本为v2。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Pool：资源池。
	Kind *string `json:"kind,omitempty"`

	Metadata *PoolMetadata `json:"metadata,omitempty"`

	Spec *PoolSpecModel `json:"spec,omitempty"`

	Status         *PoolStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolResponse", string(data)}, " ")
}
