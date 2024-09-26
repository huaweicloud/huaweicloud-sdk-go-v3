package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkErRouteTableAttachmentsResponse Response Object
type ListCentralNetworkErRouteTableAttachmentsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 创建路由表附件的返回体
	CentralNetworkErRouteTableAttachments []CentralNetworkErRouteTableAttachment `json:"central_network_er_route_table_attachments"`
	HttpStatusCode                        int                                    `json:"-"`
}

func (o ListCentralNetworkErRouteTableAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkErRouteTableAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkErRouteTableAttachmentsResponse", string(data)}, " ")
}
