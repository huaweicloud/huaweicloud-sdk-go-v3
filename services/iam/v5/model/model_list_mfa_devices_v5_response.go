package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMfaDevicesV5Response Response Object
type ListMfaDevicesV5Response struct {

	// 虚拟MFA设备列表。
	MfaDevices *[]MfaDeviceMetadata `json:"mfa_devices,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListMfaDevicesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMfaDevicesV5Response struct{}"
	}

	return strings.Join([]string{"ListMfaDevicesV5Response", string(data)}, " ")
}
