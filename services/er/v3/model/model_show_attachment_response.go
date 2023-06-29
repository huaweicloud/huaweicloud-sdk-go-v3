package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttachmentResponse Response Object
type ShowAttachmentResponse struct {
	Attachment *AttachmentResponse `json:"attachment,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttachmentResponse struct{}"
	}

	return strings.Join([]string{"ShowAttachmentResponse", string(data)}, " ")
}
