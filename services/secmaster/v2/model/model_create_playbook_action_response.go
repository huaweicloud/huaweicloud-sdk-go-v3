package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePlaybookActionResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// list of informations of playbook action
	Data *[]ActionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePlaybookActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookActionResponse struct{}"
	}

	return strings.Join([]string{"CreatePlaybookActionResponse", string(data)}, " ")
}
