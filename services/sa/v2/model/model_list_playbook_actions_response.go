package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookActionsResponse Response Object
type ListPlaybookActionsResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// current page count
	Size *int32 `json:"size,omitempty"`

	// current page size
	Page *int32 `json:"page,omitempty"`

	// list of informations of playbook action
	Data *[]ActionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookActionsResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookActionsResponse", string(data)}, " ")
}
