package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentRequest Request Object
type UpdateDocumentRequest struct {
	DocumentId string `json:"document_id"`

	Body *UpdateRunbookRequestBody `json:"body,omitempty"`
}

func (o UpdateDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDocumentRequest", string(data)}, " ")
}
