package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCentralNetworkGdgwAttachmentResponse Response Object
type ShowCentralNetworkGdgwAttachmentResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkGdgwAttachment *CentralNetworkGdgwAttachment `json:"central_network_gdgw_attachment"`
	HttpStatusCode               int                           `json:"-"`
}

func (o ShowCentralNetworkGdgwAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCentralNetworkGdgwAttachmentResponse struct{}"
	}

	return strings.Join([]string{"ShowCentralNetworkGdgwAttachmentResponse", string(data)}, " ")
}
