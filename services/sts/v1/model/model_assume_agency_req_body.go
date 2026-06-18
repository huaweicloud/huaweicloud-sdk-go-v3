package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyReqBody **参数解释**： 接口/v5/agencies/assume的Http请求体。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type AssumeAgencyReqBody struct {

	// **参数解释**： 获得的临时安全凭证的有效时间（单位：秒）。  **约束限制**： 请注意，该时间需要小于委托本身设置的最大会话持续时间，同时在携带X-Security-Token的Header头时该时间不能超过3600秒。  **取值范围**： 取值范围为[900,43200]。  **默认取值**： 默认值为3600。
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`

	// **参数解释**： 外部ID，防止混淆代理人问题。  **约束限制**： 长度范围为[2,1224]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ExternalId *string `json:"external_id,omitempty"`

	// **参数解释**： 自定义策略，限制本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。  **约束限制**： 本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。 长度范围为[2,2048]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**： 预置策略列表，限制本次会话获得的临时安全凭证的权限范围不会超过该预置策略指定的权限。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	PolicyIds *[]string `json:"policy_ids,omitempty"`

	// **参数解释**： 目标委托的URN。  **约束限制**： 长度范围为[0,1500]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	AgencyUrn string `json:"agency_urn"`

	// **参数解释**： 委托会话的会话名。  **约束限制**： 长度范围为[2,128]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	AgencySessionName string `json:"agency_session_name"`

	// **参数解释**： 调用者绑定的MFA设备的序列号。  **约束限制**： 长度范围为[9,256]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	SerialNumber *string `json:"serial_number,omitempty"`

	// **参数解释**： 调用者绑定的MFA设备上的6位数字码。  **约束限制**： 长度范围为[6,6]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	TokenCode *string `json:"token_code,omitempty"`

	// **参数解释**： 调用链里最初调用者所声明的身份。  **约束限制**： 长度范围为[2,64]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	SourceIdentity *string `json:"source_identity,omitempty"`

	// **参数解释**： 自定义标签列表。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Tags *[]TagDto `json:"tags,omitempty"`

	// **参数解释**： 随着临时安全凭证调用链持续透传的标签键列表。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	TransitiveTagKeys *[]string `json:"transitive_tag_keys,omitempty"`
}

func (o AssumeAgencyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyReqBody struct{}"
	}

	return strings.Join([]string{"AssumeAgencyReqBody", string(data)}, " ")
}
