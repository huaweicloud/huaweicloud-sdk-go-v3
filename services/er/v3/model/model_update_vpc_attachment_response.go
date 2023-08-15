package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcAttachmentResponse Response Object
type UpdateVpcAttachmentResponse struct {
	VpcAttachment *VpcAttachmentDetails `json:"vpc_attachment,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVpcAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcAttachmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcAttachmentResponse", string(data)}, " ")
}
