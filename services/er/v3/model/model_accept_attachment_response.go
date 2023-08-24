package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptAttachmentResponse Response Object
type AcceptAttachmentResponse struct {
	Attachment *AttachmentResponse `json:"attachment,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AcceptAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptAttachmentResponse struct{}"
	}

	return strings.Join([]string{"AcceptAttachmentResponse", string(data)}, " ")
}
