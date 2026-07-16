package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolNodeConfigTemplateResponse Response Object
type ShowPoolNodeConfigTemplateResponse struct {

	// **参数解释**：API版本。 **取值范围**：固定为v2。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：配置类型。 **取值范围**：固定为NodeConfigTemplate。
	Kind *string `json:"kind,omitempty"`

	Metadata *NodeconfigtemplateMetaV2 `json:"metadata,omitempty"`

	Spec *NodeconfigtemplatesSpec `json:"spec,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPoolNodeConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolNodeConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolNodeConfigTemplateResponse", string(data)}, " ")
}
