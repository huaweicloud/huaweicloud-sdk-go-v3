package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisableMfaDeviceReqBody struct {

	// IAM用户ID。
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
