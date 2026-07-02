package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPeerConfigurationSupportedDevicesResponse Response Object
type ListPeerConfigurationSupportedDevicesResponse struct {
	SupportedDevices *[]SupportedDevice `json:"supported_devices,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListPeerConfigurationSupportedDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeerConfigurationSupportedDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListPeerConfigurationSupportedDevicesResponse", string(data)}, " ")
}
