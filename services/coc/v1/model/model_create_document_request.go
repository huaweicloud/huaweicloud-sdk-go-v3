package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDocumentRequest Request Object
type CreateDocumentRequest struct {
	Body *CreateDocumentRequestBody `json:"body,omitempty"`
}

func (o CreateDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocumentRequest struct{}"
	}

	return strings.Join([]string{"CreateDocumentRequest", string(data)}, " ")
}
