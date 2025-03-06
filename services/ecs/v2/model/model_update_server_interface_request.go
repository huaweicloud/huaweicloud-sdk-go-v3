package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerInterfaceRequest Request Object
type UpdateServerInterfaceRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	// The network card ID of the cloud server.
	PortId string `json:"port_id"`

	Body *UpdateNicInfoRequestBody `json:"body,omitempty"`
}

func (o UpdateServerInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerInterfaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerInterfaceRequest", string(data)}, " ")
}
