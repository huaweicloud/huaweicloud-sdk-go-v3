package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalCocAttachmentRequest Request Object
type CreateExternalCocAttachmentRequest struct {
	Body *CreateExternalCocAttachmentRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateExternalCocAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalCocAttachmentRequest struct{}"
	}

	return strings.Join([]string{"CreateExternalCocAttachmentRequest", string(data)}, " ")
}
