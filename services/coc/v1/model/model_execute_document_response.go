package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDocumentResponse Response Object
type ExecuteDocumentResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentResponse", string(data)}, " ")
}
