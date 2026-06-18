package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InlineResponse2003OidcProvider **参数解释**： OIDC 提供商。  **取值范围**： 不涉及。
type InlineResponse2003OidcProvider struct {

	// **参数解释**： OIDC 身份提供商的 ID。  **取值范围**： 字符串长度在 1 到 64 之间，并且只能包含：字母、数字、中划线（-）。
	ProviderId string `json:"provider_id"`

	// **参数解释**： OIDC 身份提供商的名称。  **取值范围**： 字符串长度在 1 到 64 之间，并且只能包含：字母、数字、下划线（_）、中划线（-）。
	Name string `json:"name"`

	// **参数解释**： 身份提供商描述。  **取值范围**： 字符串长度不超过 255，并且不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\"、\"$\"、\"^\"和\"*\"。
	Description string `json:"description"`

	// **参数解释**： OIDC 身份提供商的 URL。  **取值范围**： 字符串长度在 1 到 255 之间。
	Url string `json:"url"`

	// **参数解释**： 统一资源名称。  **取值范围**： 字符串长度在 16 到 1500 之间，并且只能包含：字母、数字、字符\"/\"、\"=\"、\"_\"、\":\"、\"-\"。
	Urn string `json:"urn"`

	// **参数解释**： 客户端 ID 列表。  **取值范围**： 数组长度在 1 到 100 之间；数组元素为字符串，长度在 1 到 255 之间。
	ClientIds []string `json:"client_ids"`

	// **参数解释**： OIDC 身份提供商的服务器证书指纹列表。  **取值范围**： 数组长度在 1 到 5 之间；数组元素为长度为 64 的字符串，并且只能包含字母、数字。
	Thumbprints []string `json:"thumbprints"`

	// **参数解释**： 自定义标签列表。  **取值范围**： 数组长度不涉及。
	Tags []InlineResponse2001SamlProviderTags `json:"tags"`

	// **参数解释**： 提供商创建时间。  **取值范围**： 不涉及
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o InlineResponse2003OidcProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InlineResponse2003OidcProvider struct{}"
	}

	return strings.Join([]string{"InlineResponse2003OidcProvider", string(data)}, " ")
}
