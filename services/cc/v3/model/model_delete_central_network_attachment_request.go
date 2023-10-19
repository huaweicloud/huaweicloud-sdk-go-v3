package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCentralNetworkAttachmentRequest Request Object
type DeleteCentralNetworkAttachmentRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络附件ID。
	AttachmentId string `json:"attachment_id"`
}

func (o DeleteCentralNetworkAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCentralNetworkAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteCentralNetworkAttachmentRequest", string(data)}, " ")
}
