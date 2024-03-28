package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCentralNetworkErRouteTableAttachmentRequest Request Object
type ShowCentralNetworkErRouteTableAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络ER路由表附件ID。
	ErRouteTableAttachmentId string `json:"er_route_table_attachment_id"`
}

func (o ShowCentralNetworkErRouteTableAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCentralNetworkErRouteTableAttachmentRequest struct{}"
	}

	return strings.Join([]string{"ShowCentralNetworkErRouteTableAttachmentRequest", string(data)}, " ")
}
