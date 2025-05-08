package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentResponse Response Object
type UpdateDocumentResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentResponse struct{}"
	}

	return strings.Join([]string{"UpdateDocumentResponse", string(data)}, " ")
}
