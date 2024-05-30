package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerNetworkInterfaceRequest Request Object
type ChangeServerNetworkInterfaceRequest struct {

	// 网卡port id
	PortId string `json:"port_id"`

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ChangeServerNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o ChangeServerNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerNetworkInterfaceRequest", string(data)}, " ")
}
