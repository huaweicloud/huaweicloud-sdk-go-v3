package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkConnectionRequest Request Object
type UpdateCentralNetworkConnectionRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络连接ID。
	ConnectionId string `json:"connection_id"`

	Body *UpdateCentralNetworkConnectionRequestBody `json:"body,omitempty"`
}

func (o UpdateCentralNetworkConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkConnectionRequest", string(data)}, " ")
}
