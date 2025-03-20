package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisableMfaDeviceReqBody struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	// MFA设备序列号。
	SerialNumber string `json:"serial_number"`
}

func (o DisableMfaDeviceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableMfaDeviceReqBody struct{}"
	}

	return strings.Join([]string{"DisableMfaDeviceReqBody", string(data)}, " ")
}
