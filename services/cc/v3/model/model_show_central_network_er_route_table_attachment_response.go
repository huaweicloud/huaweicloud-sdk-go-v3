package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCentralNetworkErRouteTableAttachmentResponse Response Object
type ShowCentralNetworkErRouteTableAttachmentResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkErRouteTableAttachment *CentralNetworkErRouteTableAttachment `json:"central_network_er_route_table_attachment"`
	HttpStatusCode                       int                                   `json:"-"`
}

func (o ShowCentralNetworkErRouteTableAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCentralNetworkErRouteTableAttachmentResponse struct{}"
	}

	return strings.Join([]string{"ShowCentralNetworkErRouteTableAttachmentResponse", string(data)}, " ")
}
