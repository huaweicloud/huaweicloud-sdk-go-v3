package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAttachmentRequest Request Object
type UpdateAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 连接ID
	AttachmentId string `json:"attachment_id"`

	Body *UpdateAttachmentRequestBody `json:"body,omitempty"`
}

func (o UpdateAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttachmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateAttachmentRequest", string(data)}, " ")
}
