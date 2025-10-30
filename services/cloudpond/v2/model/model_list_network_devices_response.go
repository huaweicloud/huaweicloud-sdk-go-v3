package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkDevicesResponse Response Object
type ListNetworkDevicesResponse struct {

	// 网络设备列表
	NetworkDevices *[]NetworkDevice `json:"network_devices,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNetworkDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListNetworkDevicesResponse", string(data)}, " ")
}
