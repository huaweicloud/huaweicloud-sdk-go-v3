package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkDeviceResponse Response Object
type ShowNetworkDeviceResponse struct {
	NetworkDevice  *NetworkDevice `json:"network_device,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowNetworkDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowNetworkDeviceResponse", string(data)}, " ")
}
