package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPlaybookVersionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookVersionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookVersionResponse", string(data)}, " ")
}
