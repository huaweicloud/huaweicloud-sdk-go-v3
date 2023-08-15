package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookVersionResponse Response Object
type DeletePlaybookVersionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookVersionResponse", string(data)}, " ")
}
