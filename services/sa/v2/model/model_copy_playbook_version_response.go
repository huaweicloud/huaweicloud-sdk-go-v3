package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyPlaybookVersionResponse Response Object
type CopyPlaybookVersionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookVersionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyPlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyPlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"CopyPlaybookVersionResponse", string(data)}, " ")
}
