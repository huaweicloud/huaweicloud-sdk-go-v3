package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPeerConfigurationSupportedDevicesRequest Request Object
type ListPeerConfigurationSupportedDevicesRequest struct {
}

func (o ListPeerConfigurationSupportedDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeerConfigurationSupportedDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListPeerConfigurationSupportedDevicesRequest", string(data)}, " ")
}
