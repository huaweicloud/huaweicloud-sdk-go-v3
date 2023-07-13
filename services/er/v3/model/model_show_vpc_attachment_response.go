package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpcAttachmentResponse Response Object
type ShowVpcAttachmentResponse struct {
	VpcAttachment *VpcAttachmentDetails `json:"vpc_attachment,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVpcAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcAttachmentResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcAttachmentResponse", string(data)}, " ")
}
