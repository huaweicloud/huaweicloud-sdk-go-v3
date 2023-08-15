package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookResponse Response Object
type CreatePlaybookResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePlaybookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookResponse struct{}"
	}

	return strings.Join([]string{"CreatePlaybookResponse", string(data)}, " ")
}
