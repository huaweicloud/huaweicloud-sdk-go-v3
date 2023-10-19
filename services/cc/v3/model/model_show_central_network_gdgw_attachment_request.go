package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCentralNetworkGdgwAttachmentRequest Request Object
type ShowCentralNetworkGdgwAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络DGW附件ID。
	GdgwAttachmentId string `json:"gdgw_attachment_id"`
}

func (o ShowCentralNetworkGdgwAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCentralNetworkGdgwAttachmentRequest struct{}"
	}

	return strings.Join([]string{"ShowCentralNetworkGdgwAttachmentRequest", string(data)}, " ")
}
