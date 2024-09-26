package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCentralNetworkAttachmentResponse Response Object
type DeleteCentralNetworkAttachmentResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkAttachment *CentralNetworkAttachment `json:"central_network_attachment"`
	HttpStatusCode           int                       `json:"-"`
}

func (o DeleteCentralNetworkAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCentralNetworkAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteCentralNetworkAttachmentResponse", string(data)}, " ")
}
