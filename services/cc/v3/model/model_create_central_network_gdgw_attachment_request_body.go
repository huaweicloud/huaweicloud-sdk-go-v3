package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkGdgwAttachmentRequestBody 创建中心网络GDGW附件的请求体。
type CreateCentralNetworkGdgwAttachmentRequestBody struct {
	CentralNetworkGdgwAttachment *CreateCentralNetworkGdgwAttachment `json:"central_network_gdgw_attachment"`
}

func (o CreateCentralNetworkGdgwAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkGdgwAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkGdgwAttachmentRequestBody", string(data)}, " ")
}
