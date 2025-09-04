package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthenticationTemplate 创建鉴权模板请求体。
type CreateAuthenticationTemplate struct {

	// **参数说明**：鉴权模板名称。 **取值范围**：长度不超过128，只允许字母、数字、下划线（_）、连接符（-）的组合。
	TemplateName string `json:"template_name"`

	// **参数说明**：鉴权模板的描述信息。 **取值范围**：长度不超过2048，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合
	Description *string `json:"description,omitempty"`

	// **参数说明**：是否激活该鉴权模板，默认状态为未激活，只能有一个激活状态模板 - ACTIVE：该鉴权模板为激活状态。 - INACTIVE：该鉴权模板为停用状态。
	Status *string `json:"status,omitempty"`

	TemplateBody *AuthenticationTemplateBody `json:"template_body"`
}

func (o CreateAuthenticationTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthenticationTemplate struct{}"
	}

	return strings.Join([]string{"CreateAuthenticationTemplate", string(data)}, " ")
}
