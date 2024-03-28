package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkErRouteTableAttachmentResponse Response Object
type UpdateCentralNetworkErRouteTableAttachmentResponse struct {

	// 资源ID标识符。
	RequestId string `json:"request_id"`

	CentralNetworkErRouteTableAttachment *CentralNetworkErRouteTableAttachment `json:"central_network_er_route_table_attachment"`
	HttpStatusCode                       int                                   `json:"-"`
}

func (o UpdateCentralNetworkErRouteTableAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkErRouteTableAttachmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkErRouteTableAttachmentResponse", string(data)}, " ")
}
