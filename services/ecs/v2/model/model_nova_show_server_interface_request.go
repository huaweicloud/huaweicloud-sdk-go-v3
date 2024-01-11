package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaShowServerInterfaceRequest Request Object
type NovaShowServerInterfaceRequest struct {

	// 网卡port id
	PortId string `json:"port_id"`

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o NovaShowServerInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowServerInterfaceRequest struct{}"
	}

	return strings.Join([]string{"NovaShowServerInterfaceRequest", string(data)}, " ")
}
