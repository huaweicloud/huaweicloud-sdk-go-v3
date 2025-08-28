package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirtualMfaDeviceV5Request Request Object
type DeleteVirtualMfaDeviceV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	// MFA设备序列号。
	SerialNumber string `json:"serial_number"`
}

func (o DeleteVirtualMfaDeviceV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualMfaDeviceV5Request struct{}"
	}

	return strings.Join([]string{"DeleteVirtualMfaDeviceV5Request", string(data)}, " ")
}
