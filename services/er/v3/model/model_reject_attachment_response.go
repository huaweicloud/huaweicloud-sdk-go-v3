package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RejectAttachmentResponse Response Object
type RejectAttachmentResponse struct {
	Attachment *AttachmentResponse `json:"attachment,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RejectAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectAttachmentResponse struct{}"
	}

	return strings.Join([]string{"RejectAttachmentResponse", string(data)}, " ")
}
