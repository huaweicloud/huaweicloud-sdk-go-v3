package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkGdgwAttachmentResponse Response Object
type CreateCentralNetworkGdgwAttachmentResponse struct {

	// 资源ID标识符。
	RequestId string `json:"request_id"`

	CentralNetworkGdgwAttachment *CentralNetworkGdgwAttachment `json:"central_network_gdgw_attachment"`
	HttpStatusCode               int                           `json:"-"`
}

func (o CreateCentralNetworkGdgwAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkGdgwAttachmentResponse struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkGdgwAttachmentResponse", string(data)}, " ")
}
