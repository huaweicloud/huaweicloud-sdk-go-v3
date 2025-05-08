package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDocumentsResponse Response Object
type ListDocumentsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDocumentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDocumentsResponse struct{}"
	}

	return strings.Join([]string{"ListDocumentsResponse", string(data)}, " ")
}
