package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyPolicyResponseBodyPolicy **参数解释：** 密钥策略 **取值范围：** 不涉及
type UpdateKeyPolicyResponseBodyPolicy struct {

	// 密钥策略版本
	Version string `json:"version"`

	ValidityPeriod *UpdateKeyPolicyResponseBodyPolicyValidityPeriod `json:"validity_period,omitempty"`

	// 允许访问的接入点ID列表
	AllowedAccessPoint *[]string `json:"allowed_access_point,omitempty"`

	// 允许访问的数据安全专区ID列表
	AllowedDataSecurityZone *[]string `json:"allowed_data_security_zone,omitempty"`
}

func (o UpdateKeyPolicyResponseBodyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyPolicyResponseBodyPolicy struct{}"
	}

	return strings.Join([]string{"UpdateKeyPolicyResponseBodyPolicy", string(data)}, " ")
}
