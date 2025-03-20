package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsymmetricSignatureWithDomainId 账号非对称签名开关信息。
type AsymmetricSignatureWithDomainId struct {

	// 账号ID。
	DomainId string `json:"domain_id"`

	// 凭证非对称签名开关。
	AsymmetricSignatureSwitch bool `json:"asymmetric_signature_switch"`
}

func (o AsymmetricSignatureWithDomainId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsymmetricSignatureWithDomainId struct{}"
	}

	return strings.Join([]string{"AsymmetricSignatureWithDomainId", string(data)}, " ")
}
