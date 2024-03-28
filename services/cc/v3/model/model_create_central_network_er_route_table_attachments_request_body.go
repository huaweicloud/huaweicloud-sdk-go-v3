package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkErRouteTableAttachmentsRequestBody 创建中心网络ER路由表附件的请求体。
type CreateCentralNetworkErRouteTableAttachmentsRequestBody struct {
	CentralNetworkErRouteTableAttachment *CreateCentralNetworkErRouteTableAttachment `json:"central_network_er_route_table_attachment"`
}

func (o CreateCentralNetworkErRouteTableAttachmentsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkErRouteTableAttachmentsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkErRouteTableAttachmentsRequestBody", string(data)}, " ")
}
