package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookVersionsResponse Response Object
type ListPlaybookVersionsResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// current page count
	Size *int32 `json:"size,omitempty"`

	// current page size
	Page *int32 `json:"page,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// list of informations of playbook version
	Data *[]PlaybookVersionListEntity `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookVersionsResponse", string(data)}, " ")
}
