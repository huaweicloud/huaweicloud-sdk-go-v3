package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableMfaDeviceReqBody struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	// MFA设备序列号。
	SerialNumber string `json:"serial_number"`

	// 设备发出的验证码。
	AuthenticationCodeFirst string `json:"authentication_code_first"`

	// 设备发出的后续验证码。
	AuthenticationCodeSecond string `json:"authentication_code_second"`
}

func (o EnableMfaDeviceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableMfaDeviceReqBody struct{}"
	}

	return strings.Join([]string{"EnableMfaDeviceReqBody", string(data)}, " ")
}
