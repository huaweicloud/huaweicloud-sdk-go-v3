package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSubNetworkInterfaceRequest struct {
	// 辅助弹性网卡的唯一标识

	SubNetworkInterfaceId string `json:"sub_network_interface_id"`
}

func (o ShowSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfaceRequest", string(data)}, " ")
}
