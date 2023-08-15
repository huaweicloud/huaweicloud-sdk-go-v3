package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookActionResponse Response Object
type DeletePlaybookActionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePlaybookActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookActionResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookActionResponse", string(data)}, " ")
}
