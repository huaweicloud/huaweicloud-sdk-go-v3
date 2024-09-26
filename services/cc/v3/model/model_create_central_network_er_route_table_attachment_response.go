package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkErRouteTableAttachmentResponse Response Object
type CreateCentralNetworkErRouteTableAttachmentResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkErRouteTableAttachment *CentralNetworkErRouteTableAttachment `json:"central_network_er_route_table_attachment"`
	HttpStatusCode                       int                                   `json:"-"`
}

func (o CreateCentralNetworkErRouteTableAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkErRouteTableAttachmentResponse struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkErRouteTableAttachmentResponse", string(data)}, " ")
}
