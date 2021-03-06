package model

import (
	"encoding/json"

	"strings"
)

// MFA设备密钥。
type CreateMfaDeviceRespon struct {
	// MFA设备序列号。

	SerialNumber string `json:"serial_number"`
	// 密钥信息，用于第三方生成图片验证码。

	Base32StringSeed string `json:"base32_string_seed"`
}

func (o CreateMfaDeviceRespon) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMfaDeviceRespon struct{}"
	}

	return strings.Join([]string{"CreateMfaDeviceRespon", string(data)}, " ")
}
