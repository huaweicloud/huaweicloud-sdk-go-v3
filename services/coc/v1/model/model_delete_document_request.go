package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDocumentRequest Request Object
type DeleteDocumentRequest struct {

	// 作业uuid
	DocumentId string `json:"document_id"`
}

func (o DeleteDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDocumentRequest struct{}"
	}

	return strings.Join([]string{"DeleteDocumentRequest", string(data)}, " ")
}
