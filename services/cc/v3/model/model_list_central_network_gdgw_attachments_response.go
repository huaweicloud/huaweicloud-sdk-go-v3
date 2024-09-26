package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkGdgwAttachmentsResponse Response Object
type ListCentralNetworkGdgwAttachmentsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 中心网络GDGW附件列表。
	CentralNetworkGdgwAttachments []CentralNetworkGdgwAttachment `json:"central_network_gdgw_attachments"`
	HttpStatusCode                int                            `json:"-"`
}

func (o ListCentralNetworkGdgwAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkGdgwAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkGdgwAttachmentsResponse", string(data)}, " ")
}
