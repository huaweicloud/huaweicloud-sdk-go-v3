package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MfaDeviceDto struct {

	// MFA设备唯一标识符（ID）
	DeviceId string `json:"device_id"`

	// MFA设备名称
	DeviceName string `json:"device_name"`

	// MFA设备显示名称
	DisplayName string `json:"display_name"`

	// MFA类型
	MfaType string `json:"mfa_type"`

	// 注册时间
	RegisteredDate int64 `json:"registered_date"`
}

func (o MfaDeviceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaDeviceDto struct{}"
	}

	return strings.Join([]string{"MfaDeviceDto", string(data)}, " ")
}
