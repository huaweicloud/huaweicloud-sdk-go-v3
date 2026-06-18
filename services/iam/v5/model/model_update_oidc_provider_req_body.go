package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOidcProviderReqBody **参数解释**： 更新 OIDC 提供商的请求体。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type UpdateOidcProviderReqBody struct {

	// **参数解释**： 身份提供商描述。  **约束限制**： 长度范围为[0,255]。  **取值范围**： 不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\"、\"$\"、\"^\"和\"*\"的字符串。  **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o UpdateOidcProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOidcProviderReqBody struct{}"
	}

	return strings.Join([]string{"UpdateOidcProviderReqBody", string(data)}, " ")
}
