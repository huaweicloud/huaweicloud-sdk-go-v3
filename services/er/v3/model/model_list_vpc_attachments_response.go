package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcAttachmentsResponse Response Object
type ListVpcAttachmentsResponse struct {

	// VPC连接列表
	VpcAttachments *[]VpcAttachmentDetails `json:"vpc_attachments,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVpcAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcAttachmentsResponse", string(data)}, " ")
}
