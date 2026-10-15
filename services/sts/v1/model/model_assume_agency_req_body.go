package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyReqBody **参数解释**： 接口/v5/agencies/assume的Http请求体。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type AssumeAgencyReqBody struct {

	// **参数解释**： 获得的临时安全凭证的有效时间（单位：秒）。  **约束限制**： 请注意，该时间需要小于委托本身设置的最大会话持续时间（在IAM控制台「委托」页面中获取），同时在携带X-Security-Token的Header头时该时间不能超过3600秒。  **取值范围**： 取值范围为[900,43200]。  **默认取值**： 默认值为3600。
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`

	// **参数解释**： 外部ID，防止混淆代理人问题。  **约束限制**： 不涉及。  **取值范围**： 长度为2到1224个字符，只允许大小写字母、数字、\"_\"、\"+\"、\"=\"、\",\"、\".\"、\"@\"、\":\"、\"/\"、\"-\"。  **默认取值**： 不涉及。
	ExternalId *string `json:"external_id,omitempty"`

	// **参数解释**： 自定义策略，用于进一步限制临时安全凭证的权限范围。  **约束限制**： 本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。  **取值范围**： 长度为2到4096个字符。  **默认取值**： 不涉及。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**： 预置策略列表，限制本次会话获得的临时安全凭证的权限范围不会超过该预置策略指定的权限。 每个元素为预置策略的标识ID，长度限制为1到64个字符，只允许大小写字母、数字、\"_\"、\"-\"，且至少包含一个字符。  **约束限制**： 该数组允许的元素个数最多为20。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	PolicyIds *[]string `json:"policy_ids,omitempty"`

	// **参数解释**： 目标委托的URN，在IAM控制台的「委托」页面中获取。  **约束限制**： 不涉及。  **取值范围**： 长度最多为1500个字符。格式为5段，用\":\"分隔，各段允许大小写字母、数字、\"+\"、\"/\"、\"=\"、\"-\"、\"_\"（第2段还额外允许 \"*\"），长度限制为：第1段至少1位，第2段0~255位，第3段1~64位，第4段1~64位，第5段至少1位（任意字符）。  **默认取值**： 不涉及。
	AgencyUrn string `json:"agency_urn"`

	// **参数解释**： 委托会话的会话名。  **约束限制**： 不涉及。  **取值范围**： 长度为2到128个字符，只允许大小写字母、数字、\"_\"、\"+\"、\"=\"、\",\"、\".\"、\"@\"、\"-\"。  **默认取值**： 不涉及。
	AgencySessionName string `json:"agency_session_name"`

	// **参数解释**： 调用者绑定的MFA设备的序列号。在IAM控制台->「用户」->「安全设置」->「多因素认证设备」中获取。  **约束限制**： 不涉及。  **取值范围**： 长度为9到256个字符，只允许大小写字母、数字、\"_\"、\"+\"、\"=\"、\"/\"、\":\"、\",\"、\".\"、\"@\"、\"-\"。  **默认取值**： 不涉及。
	SerialNumber *string `json:"serial_number,omitempty"`

	// **参数解释**： 调用者绑定的MFA设备上的6位纯数字。  **约束限制**： 只允许6位数字。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	TokenCode *string `json:"token_code,omitempty"`

	// **参数解释**： 调用链里最初调用者所声明的身份。  **约束限制**： 不涉及。  **取值范围**： 长度为2到64个字符，只允许大小写字母、数字、\"_\"、\"+\"、\"=\"、\",\"、\".\"、\"@\"、\"-\"。  **默认取值**： 不涉及。
	SourceIdentity *string `json:"source_identity,omitempty"`

	// **参数解释**： 自定义标签列表。  **约束限制**： 该数组允许的元素个数最多为20。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Tags *[]TagDto `json:"tags,omitempty"`

	// **参数解释**： 随着临时安全凭证调用链持续透传的标签键列表。  **约束限制**： 该数组允许的元素个数最多为20。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	TransitiveTagKeys *[]string `json:"transitive_tag_keys,omitempty"`

	// **参数解释**： 预先获取的可信上下文断言列表，格式为数组。  **约束限制**： 该数组允许的元素个数为1到5。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ProvidedContexts *[]ProvidedContextDto `json:"provided_contexts,omitempty"`
}

func (o AssumeAgencyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyReqBody struct{}"
	}

	return strings.Join([]string{"AssumeAgencyReqBody", string(data)}, " ")
}
