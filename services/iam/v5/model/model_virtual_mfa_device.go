package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VirtualMfaDevice MFA设备。
type VirtualMfaDevice struct {

	// MFA设备序列号。
	SerialNumber string `json:"serial_number"`

	// 密钥信息，用于第三方生成图片验证码。
	Base32StringSeed string `json:"base32_string_seed"`
}

func (o VirtualMfaDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualMfaDevice struct{}"
	}

	return strings.Join([]string{"VirtualMfaDevice", string(data)}, " ")
}
