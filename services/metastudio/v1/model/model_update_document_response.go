package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentResponse Response Object
type UpdateDocumentResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentResponse struct{}"
	}

	return strings.Join([]string{"UpdateDocumentResponse", string(data)}, " ")
}
