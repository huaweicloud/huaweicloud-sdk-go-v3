package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkErRouteTableAttachmentRequest Request Object
type UpdateCentralNetworkErRouteTableAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络ER路由表附件ID。
	ErRouteTableAttachmentId string `json:"er_route_table_attachment_id"`

	Body *UpdateCentralNetworkErRouteTableAttachmentRequestBody `json:"body,omitempty"`
}

func (o UpdateCentralNetworkErRouteTableAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkErRouteTableAttachmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkErRouteTableAttachmentRequest", string(data)}, " ")
}
