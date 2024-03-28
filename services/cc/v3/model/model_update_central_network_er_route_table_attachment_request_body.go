package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkErRouteTableAttachmentRequestBody 更新中心网络ER附件的请求体。
type UpdateCentralNetworkErRouteTableAttachmentRequestBody struct {
	CentralNetworkErRouteTableAttachment *UpdateCentralNetworkErRouteTableAttachment `json:"central_network_er_route_table_attachment"`
}

func (o UpdateCentralNetworkErRouteTableAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkErRouteTableAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkErRouteTableAttachmentRequestBody", string(data)}, " ")
}
