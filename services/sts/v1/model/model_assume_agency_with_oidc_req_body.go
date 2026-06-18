package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyWithOidcReqBody **参数解释**： 接口/v5/agencies/assume-with-oidc的Http请求体。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type AssumeAgencyWithOidcReqBody struct {

	// **参数解释**： 获得的临时安全凭证的有效时间（单位：秒）。  **约束限制**： 请注意，该时间需要小于委托本身设置的最大会话持续时间，同时在携带X-Security-Token的Header头时该时间不能超过3600秒。  **取值范围**： 取值范围为[900,43200]。  **默认取值**： 默认值为3600。
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`

	// **参数解释**： 自定义策略，限制本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。  **约束限制**： 本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。 长度范围为[2,2048]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**： 预置策略列表，限制本次会话获得的临时安全凭证的权限范围不会超过该预置策略指定的权限。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	PolicyIds *[]string `json:"policy_ids,omitempty"`

	// **参数解释**： OIDC提供商的URN。  **约束限制**： 长度范围为[0,1500]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ProviderUrn string `json:"provider_urn"`

	// **参数解释**： 目标信任委托的URN。  **约束限制**： 长度范围为[0,1500]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	AgencyUrn string `json:"agency_urn"`

	// **参数解释**： 信任委托会话的会话名。  **约束限制**： 长度范围为[2,128]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	AgencySessionName string `json:"agency_session_name"`

	// **参数解释**： 由身份提供商提供的OIDC令牌。  **约束限制**： 长度范围为[4,20000]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	IdToken string `json:"id_token"`
}

func (o AssumeAgencyWithOidcReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyWithOidcReqBody struct{}"
	}

	return strings.Join([]string{"AssumeAgencyWithOidcReqBody", string(data)}, " ")
}
