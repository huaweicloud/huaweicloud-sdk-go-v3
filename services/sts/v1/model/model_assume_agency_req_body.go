package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyReqBody 接口/v5/agencies/assume的Http请求体。
type AssumeAgencyReqBody struct {

	// 获得的临时安全凭证的有效时间（单位秒）。请注意，该时间需要小于委托本身设置的最大会话持续时间，同时在携带X-Security-Token的Header头时该时间不能超过3600秒。
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`

	// 外部ID，防止混淆代理人问题。
	ExternalId *string `json:"external_id,omitempty"`

	// 自定义策略，限制本次会话获得的临时安全凭证的权限范围不会超过该自定义策略指定的权限。
	Policy *string `json:"policy,omitempty"`

	// 预置策略列表，限制本次会话获得的临时安全凭证的权限范围不会超过该预置策略指定的权限。
	PolicyIds *[]string `json:"policy_ids,omitempty"`

	// 目标委托的URN。
	AgencyUrn string `json:"agency_urn"`

	// 委托会话的会话名。
	AgencySessionName string `json:"agency_session_name"`

	// 调用者绑定的MFA设备的序列号。
	SerialNumber *string `json:"serial_number,omitempty"`

	// 调用者绑定的MFA设备上的6位数字码。
	TokenCode *string `json:"token_code,omitempty"`

	// 调用链里最初调用者所声明的身份。
	SourceIdentity *string `json:"source_identity,omitempty"`

	// 自定义标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`

	// 随着临时安全凭证调用链持续透传的标签键列表。
	TransitiveTagKeys *[]string `json:"transitive_tag_keys,omitempty"`
}

func (o AssumeAgencyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyReqBody struct{}"
	}

	return strings.Join([]string{"AssumeAgencyReqBody", string(data)}, " ")
}
