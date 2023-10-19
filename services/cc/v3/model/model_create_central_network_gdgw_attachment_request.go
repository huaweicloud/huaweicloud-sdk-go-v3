package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkGdgwAttachmentRequest Request Object
type CreateCentralNetworkGdgwAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	Body *CreateCentralNetworkGdgwAttachmentRequestBody `json:"body,omitempty"`
}

func (o CreateCentralNetworkGdgwAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkGdgwAttachmentRequest struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkGdgwAttachmentRequest", string(data)}, " ")
}
