package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcAttachmentRequest Request Object
type UpdateVpcAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// VPC连接ID
	VpcAttachmentId string `json:"vpc_attachment_id"`

	Body *UpdateVpcAttachmentRequestBody `json:"body,omitempty"`
}

func (o UpdateVpcAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcAttachmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpcAttachmentRequest", string(data)}, " ")
}
