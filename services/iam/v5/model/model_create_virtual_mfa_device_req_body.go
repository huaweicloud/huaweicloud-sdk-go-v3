package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVirtualMfaDeviceReqBody struct {

	// MFA设备名称，长度为1到64个字符，只包含字母、数字、\"_\"和\"-\"的字符串。
	VirtualMfaDeviceName string `json:"virtual_mfa_device_name"`

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`
}

func (o CreateVirtualMfaDeviceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualMfaDeviceReqBody struct{}"
	}

	return strings.Join([]string{"CreateVirtualMfaDeviceReqBody", string(data)}, " ")
}
