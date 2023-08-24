package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptAttachmentRequest Request Object
type AcceptAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 连接ID
	AttachmentId string `json:"attachment_id"`
}

func (o AcceptAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptAttachmentRequest struct{}"
	}

	return strings.Join([]string{"AcceptAttachmentRequest", string(data)}, " ")
}
