package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDocumentResponse Response Object
type DeleteDocumentResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDocumentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDocumentResponse", string(data)}, " ")
}
