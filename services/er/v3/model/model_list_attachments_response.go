package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAttachmentsResponse struct {

	// 连接列表
	Attachments *[]AttachmentDetails `json:"attachments,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListAttachmentsResponse", string(data)}, " ")
}
