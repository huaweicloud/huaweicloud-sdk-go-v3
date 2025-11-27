package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDocumentRequest Request Object
type ExecuteDocumentRequest struct {
	DocumentId string `json:"document_id"`

	Body *ExecuteDocumentRequestBody `json:"body,omitempty"`
}

func (o ExecuteDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequest", string(data)}, " ")
}
