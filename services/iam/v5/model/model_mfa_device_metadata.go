package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MfaDeviceMetadata 虚拟MFA设备。
type MfaDeviceMetadata struct {

	// MFA设备序列号。
	SerialNumber string `json:"serial_number"`

	// IAM用户ID。
	UserId string `json:"user_id"`

	// 虚拟MFA设备是否开启。
	Enabled bool `json:"enabled"`
}

func (o MfaDeviceMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaDeviceMetadata struct{}"
	}

	return strings.Join([]string{"MfaDeviceMetadata", string(data)}, " ")
}
