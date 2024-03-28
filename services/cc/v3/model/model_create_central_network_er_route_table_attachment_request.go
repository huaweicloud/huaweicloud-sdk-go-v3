package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkErRouteTableAttachmentRequest Request Object
type CreateCentralNetworkErRouteTableAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	Body *CreateCentralNetworkErRouteTableAttachmentsRequestBody `json:"body,omitempty"`
}

func (o CreateCentralNetworkErRouteTableAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkErRouteTableAttachmentRequest struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkErRouteTableAttachmentRequest", string(data)}, " ")
}
