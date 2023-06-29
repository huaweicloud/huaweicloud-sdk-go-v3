package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookResponse Response Object
type ShowPlaybookResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookResponse", string(data)}, " ")
}
