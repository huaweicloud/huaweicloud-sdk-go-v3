package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkGdgwAttachmentResponse Response Object
type UpdateCentralNetworkGdgwAttachmentResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkGdgwAttachment *CentralNetworkGdgwAttachment `json:"central_network_gdgw_attachment"`
	HttpStatusCode               int                           `json:"-"`
}

func (o UpdateCentralNetworkGdgwAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkGdgwAttachmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkGdgwAttachmentResponse", string(data)}, " ")
}
