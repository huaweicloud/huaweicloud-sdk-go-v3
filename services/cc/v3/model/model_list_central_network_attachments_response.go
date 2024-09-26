package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkAttachmentsResponse Response Object
type ListCentralNetworkAttachmentsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 接入网络实例连接列表。
	CentralNetworkAttachments []CentralNetworkAttachment `json:"central_network_attachments"`
	HttpStatusCode            int                        `json:"-"`
}

func (o ListCentralNetworkAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkAttachmentsResponse", string(data)}, " ")
}
