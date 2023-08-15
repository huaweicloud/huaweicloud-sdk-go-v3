package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookVersionResponse Response Object
type UpdatePlaybookVersionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookVersionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookVersionResponse", string(data)}, " ")
}
