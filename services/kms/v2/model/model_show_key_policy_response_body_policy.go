package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeyPolicyResponseBodyPolicy **参数解释：** 密钥策略 **取值范围：** 不涉及
type ShowKeyPolicyResponseBodyPolicy struct {

	// **参数解释：** 密钥策略版本 **取值范围：** 不涉及
	Version string `json:"version"`

	ValidityPeriod *ListKeyPolicyResponseBodyPolicyValidityPeriod `json:"validity_period,omitempty"`

	// **参数解释：** 允许访问的接入点ID列表 **取值范围：** 不涉及
	AllowedAccessPoint *[]string `json:"allowed_access_point,omitempty"`

	// **参数解释：** 允许访问的数据安全专区ID列表 **取值范围：** 不涉及
	AllowedDataSecurityZone *[]string `json:"allowed_data_security_zone,omitempty"`
}

func (o ShowKeyPolicyResponseBodyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyPolicyResponseBodyPolicy struct{}"
	}

	return strings.Join([]string{"ShowKeyPolicyResponseBodyPolicy", string(data)}, " ")
}
