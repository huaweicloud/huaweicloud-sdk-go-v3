package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcAttachmentRequestBody This is a auto create Body Object
type UpdateVpcAttachmentRequestBody struct {
	VpcAttachment *UpdateVpcAttachmentBody `json:"vpc_attachment,omitempty"`
}

func (o UpdateVpcAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcAttachmentRequestBody", string(data)}, " ")
}
