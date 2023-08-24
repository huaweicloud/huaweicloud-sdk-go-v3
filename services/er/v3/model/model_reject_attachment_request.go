package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RejectAttachmentRequest Request Object
type RejectAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 连接ID
	AttachmentId string `json:"attachment_id"`
}

func (o RejectAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectAttachmentRequest struct{}"
	}

	return strings.Join([]string{"RejectAttachmentRequest", string(data)}, " ")
}
