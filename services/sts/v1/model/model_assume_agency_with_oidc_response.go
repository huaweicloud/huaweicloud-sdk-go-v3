package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyWithOidcResponse Response Object
type AssumeAgencyWithOidcResponse struct {

	// **参数解释**： 身份提供商返回的OIDC令牌中所申明的身份。  **取值范围**： 不涉及。
	SourceIdentity *string `json:"source_identity,omitempty"`

	AssumedAgency *AssumedAgencyWithFederationDto `json:"assumed_agency,omitempty"`

	Credentials *CredentialsDto `json:"credentials,omitempty"`

	// **参数解释**： OIDC令牌的预期受众（也称为客户端ID），通常是分发给应用程序的客户端标识符。  **取值范围**： 不涉及。
	Audience *string `json:"audience,omitempty"`

	// **参数解释**： 身份提供商的URN。  **取值范围**： 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释**： 由身份提供商返回的唯一用户标识符，即OIDC令牌中的sub(Subject)声明的值。  **取值范围**： 不涉及。
	SubjectFromIdToken *string `json:"subject_from_id_token,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o AssumeAgencyWithOidcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyWithOidcResponse struct{}"
	}

	return strings.Join([]string{"AssumeAgencyWithOidcResponse", string(data)}, " ")
}
