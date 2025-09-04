package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthenticationTemplate 更新鉴权模板请求体。
type UpdateAuthenticationTemplate struct {

	// **参数说明**：鉴权模板的描述信息。 **取值范围**：长度不超过2048，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合
	Description *string `json:"description,omitempty"`

	// **参数说明**：是否激活该鉴权模板 - ACTIVE：该鉴权模板为激活状态。 - INACTIVE：该鉴权模板为停用状态。
	Status *string `json:"status,omitempty"`

	TemplateBody *UpdateAuthenticationTemplateBody `json:"template_body,omitempty"`
}

func (o UpdateAuthenticationTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthenticationTemplate struct{}"
	}

	return strings.Join([]string{"UpdateAuthenticationTemplate", string(data)}, " ")
}
