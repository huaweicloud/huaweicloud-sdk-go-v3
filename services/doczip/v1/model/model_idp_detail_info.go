package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdpDetailInfo 身份提供方信息
type IdpDetailInfo struct {

	// **参数解释**： 身份提供方类型。 **约束限制**： 不涉及。 **取值范围**： - oidc：OIDC签名。 - self_signed_jwt：自签。 **默认取值**： 不涉及。
	IdpType *string `json:"idp_type,omitempty"`

	// **参数解释**： 身份提供方的 Issuer URL。 **约束限制**： 格式如：https://idp.example.com。 **取值范围**： [1-512]个字符。 **默认取值**： 不涉及。
	IdpIssuer *string `json:"idp_issuer,omitempty"`

	// **参数解释**： 签名算法。 **约束限制**： 不涉及。 **取值范围**： - RS256：RS256算法。 - ES256：ES256算法。 **默认取值**： 不涉及。
	IdpSigningAlg *string `json:"idp_signing_alg,omitempty"`

	// **参数解释**： 预期的 Audience。 **约束限制**： 不涉及。 **取值范围**： [1-512]个字符。 **默认取值**： 不涉及。
	IdpAudience *string `json:"idp_audience,omitempty"`

	// **参数解释**： 身份提供方公钥。 **取值范围**： 不涉及。
	PublicKeys *[]PublicInfo `json:"public_keys,omitempty"`

	// **参数解释**： [创建时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [创建时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// **参数解释**： [创建时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [创建时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o IdpDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpDetailInfo struct{}"
	}

	return strings.Join([]string{"IdpDetailInfo", string(data)}, " ")
}
