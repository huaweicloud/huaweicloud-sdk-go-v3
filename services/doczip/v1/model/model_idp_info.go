package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdpInfo 身份提供方更新信息
type IdpInfo struct {

	// **参数解释**： 身份提供方类型。 **约束限制**： 不涉及。 **取值范围**： - oidc：OIDC签名。 - self_signed_jwt：自签。 **默认取值**： 不涉及。
	IdpType string `json:"idp_type"`

	// **参数解释**： 身份提供方的 Issuer URL。 **约束限制**： 格式如：https://idp.example.com。 **取值范围**： [1-512]个字符。 **默认取值**： 不涉及。
	IdpIssuer *string `json:"idp_issuer,omitempty"`

	// **参数解释**： 签名算法。 **约束限制**： 不涉及。 **取值范围**： - RS256：RS256算法。 - ES256：ES256算法。 **默认取值**： 不涉及。
	IdpSigningAlg *string `json:"idp_signing_alg,omitempty"`

	// **参数解释**： 预期的 Audience。 **约束限制**： 不涉及。 **取值范围**： [1-512]个字符。 **默认取值**： 不涉及。
	IdpAudience *string `json:"idp_audience,omitempty"`

	PublicKeys *PublicUpdateInfo `json:"public_keys,omitempty"`
}

func (o IdpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpInfo struct{}"
	}

	return strings.Join([]string{"IdpInfo", string(data)}, " ")
}
