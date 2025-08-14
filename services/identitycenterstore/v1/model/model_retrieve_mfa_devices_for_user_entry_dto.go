package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetrieveMfaDevicesForUserEntryDto struct {

	// MFA设备列表
	MfaDevices []MfaDeviceDto `json:"mfa_devices"`

	User *RetrieveMfaDevicesForUserDto `json:"user"`
}

func (o RetrieveMfaDevicesForUserEntryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrieveMfaDevicesForUserEntryDto struct{}"
	}

	return strings.Join([]string{"RetrieveMfaDevicesForUserEntryDto", string(data)}, " ")
}
