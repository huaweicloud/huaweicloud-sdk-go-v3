package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsymmetricSignature 设置账号是否开启非对称签名功能。
type AsymmetricSignature struct {

	// 非对称签名开关。
	AsymmetricSignatureSwitch bool `json:"asymmetric_signature_switch"`
}

func (o AsymmetricSignature) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsymmetricSignature struct{}"
	}

	return strings.Join([]string{"AsymmetricSignature", string(data)}, " ")
}
