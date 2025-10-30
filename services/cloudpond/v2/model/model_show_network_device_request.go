package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkDeviceRequest Request Object
type ShowNetworkDeviceRequest struct {

	// 网络设备ID
	NetworkDeviceId string `json:"network_device_id"`
}

func (o ShowNetworkDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkDeviceRequest struct{}"
	}

	return strings.Join([]string{"ShowNetworkDeviceRequest", string(data)}, " ")
}
