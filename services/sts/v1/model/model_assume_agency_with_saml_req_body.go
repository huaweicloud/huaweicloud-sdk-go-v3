package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyWithSamlReqBody **参数解释**： 接口/v5/agencies/assume-with-saml的Http请求体。  **取值范围**： 不涉及。
type AssumeAgencyWithSamlReqBody struct {

	// **参数解释**： 获得的临时安全凭证的有效时间（单位：秒）。  **约束限制**： 获得的临时安全凭证的有效时间（单位：秒）。请注意，该时间需要小于信任委托本身设置的最大会话持续时间。同时最终的会话持续时间以duration_seconds，  SAML身份验证响应中SessionNotOnOrAfter值和SessionDuration值三者中较短的一个为准。  **取值范围**： 取值范围为[900,43200]。  **默认取值**： 默认值为3600。
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`

	// **参数解释**： 自定义策略，限制本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。  **约束限制**： 本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。 长度范围为[2,2048]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**： 预置策略列表，限制本次会话获得的临时安全凭证的权限范围不会超过该预置策略指定的权限。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	PolicyIds *[]string `json:"policy_ids,omitempty"`

	// **参数解释**： SAML提供商的URN。  **约束限制**： 长度范围为[0,1500]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ProviderUrn string `json:"provider_urn"`

	// **参数解释**： 目标信任委托的URN。  **约束限制**： 长度范围为[0,1500]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	AgencyUrn string `json:"agency_urn"`

	// **参数解释**： 由SAML身份提供商提供的Base64编码的SAML身份验证响应。  **约束限制**： 长度范围为[4,100000]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	SamlAssertion string `json:"saml_assertion"`
}

func (o AssumeAgencyWithSamlReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyWithSamlReqBody struct{}"
	}

	return strings.Join([]string{"AssumeAgencyWithSamlReqBody", string(data)}, " ")
}
