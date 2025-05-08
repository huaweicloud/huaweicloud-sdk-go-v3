package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDocumentResponse Response Object
type CreateDocumentResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocumentResponse struct{}"
	}

	return strings.Join([]string{"CreateDocumentResponse", string(data)}, " ")
}
