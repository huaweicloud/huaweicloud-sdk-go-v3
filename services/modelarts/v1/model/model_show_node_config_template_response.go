package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeConfigTemplateResponse Response Object
type ShowNodeConfigTemplateResponse struct {

	// **参数解释**： API版本。 **取值范围**： 可选值如下： - v1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 资源类型。 **取值范围**： 可选值如下： - NodeConfigTemplate：节点配置模板
	Kind *string `json:"kind,omitempty"`

	Metadata *NodeConfigTemplateMeta `json:"metadata,omitempty"`

	Spec           *NodeConfigTemplateSpec `json:"spec,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowNodeConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeConfigTemplateResponse", string(data)}, " ")
}
