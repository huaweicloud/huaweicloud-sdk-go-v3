package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkGdgwAttachmentRequestBody 更新中心网络GDGW附件的请求体。
type UpdateCentralNetworkGdgwAttachmentRequestBody struct {
	CentralNetworkGdgwAttachment *UpdateCentralNetworkGdgwAttachment `json:"central_network_gdgw_attachment"`
}

func (o UpdateCentralNetworkGdgwAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkGdgwAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkGdgwAttachmentRequestBody", string(data)}, " ")
}
