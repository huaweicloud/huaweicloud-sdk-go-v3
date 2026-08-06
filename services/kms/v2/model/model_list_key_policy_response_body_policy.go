package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeyPolicyResponseBodyPolicy **参数解释：** 密钥策略 **取值范围：** 不涉及
type ListKeyPolicyResponseBodyPolicy struct {

	// **参数解释：** 密钥策略版本 **取值范围：** 不涉及
	Version string `json:"version"`

	ValidityPeriod *ListKeyPolicyResponseBodyPolicyValidityPeriod `json:"validityPeriod,omitempty"`

	// **参数解释：** 允许访问的接入点ID列表 **取值范围：** 不涉及
	AllowedAccessPoint *[]string `json:"allowedAccessPoint,omitempty"`

	// **参数解释：** 允许访问的数据安全专区ID列表 **取值范围：** 不涉及
	AllowedDataSecurityZone *[]string `json:"allowed_data_security_zone,omitempty"`
}

func (o ListKeyPolicyResponseBodyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyPolicyResponseBodyPolicy struct{}"
	}

	return strings.Join([]string{"ListKeyPolicyResponseBodyPolicy", string(data)}, " ")
}
