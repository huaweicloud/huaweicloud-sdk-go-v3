package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolNodeConfigResponse Response Object
type ShowPoolNodeConfigResponse struct {

	// **参数解释**： 固定为v2。 **取值范围**： 不涉及。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 固定为NodeConfig。 **取值范围**： 不涉及。
	Kind *string `json:"kind,omitempty"`

	Metadata *NodeconfigMeta `json:"metadata,omitempty"`

	Spec *NodeconfigSpec `json:"spec,omitempty"`

	Status *NodeconfigStatus `json:"status,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPoolNodeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolNodeConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolNodeConfigResponse", string(data)}, " ")
}
