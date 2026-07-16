package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeConfigTemplateRequest Request Object
type ShowNodeConfigTemplateRequest struct {

	// **参数解释**：节点配置模板的名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodeconfigtemplateName string `json:"nodeconfigtemplate_name"`
}

func (o ShowNodeConfigTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeConfigTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowNodeConfigTemplateRequest", string(data)}, " ")
}
