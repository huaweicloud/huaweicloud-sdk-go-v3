package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttachmentRequest Request Object
type ShowAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 连接ID
	AttachmentId string `json:"attachment_id"`
}

func (o ShowAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttachmentRequest struct{}"
	}

	return strings.Join([]string{"ShowAttachmentRequest", string(data)}, " ")
}
