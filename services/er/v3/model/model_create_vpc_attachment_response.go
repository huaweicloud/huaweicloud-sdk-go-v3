package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVpcAttachmentResponse struct {
	VpcAttachment *VpcAttachmentDetails `json:"vpc_attachment,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVpcAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcAttachmentResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcAttachmentResponse", string(data)}, " ")
}
