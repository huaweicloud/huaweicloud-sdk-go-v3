package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDocumentResponse Response Object
type DeleteDocumentResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDocumentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDocumentResponse", string(data)}, " ")
}
