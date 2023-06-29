package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAttachmentResponse Response Object
type UpdateAttachmentResponse struct {
	Attachment *AttachmentResponse `json:"attachment,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttachmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateAttachmentResponse", string(data)}, " ")
}
