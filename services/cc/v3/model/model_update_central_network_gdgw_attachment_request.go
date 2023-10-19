package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkGdgwAttachmentRequest Request Object
type UpdateCentralNetworkGdgwAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络DGW附件ID。
	GdgwAttachmentId string `json:"gdgw_attachment_id"`

	Body *UpdateCentralNetworkGdgwAttachmentRequestBody `json:"body,omitempty"`
}

func (o UpdateCentralNetworkGdgwAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkGdgwAttachmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkGdgwAttachmentRequest", string(data)}, " ")
}
