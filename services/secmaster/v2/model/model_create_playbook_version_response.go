package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookVersionResponse Response Object
type CreatePlaybookVersionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookVersionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"CreatePlaybookVersionResponse", string(data)}, " ")
}
