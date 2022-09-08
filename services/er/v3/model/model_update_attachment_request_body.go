package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAttachmentRequestBody struct {
	Attachment *UpdateAttachmentBody `json:"attachment,omitempty"`
}

func (o UpdateAttachmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttachmentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAttachmentRequestBody", string(data)}, " ")
}
