package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentTypesResponse Response Object
type ListIncidentTypesResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// Response of dataclass detail
	Data *[]DataClassTypeDetailInfo `json:"data,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// current page size
	Size *int32 `json:"size,omitempty"`

	// current page count
	Page *int32 `json:"page,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIncidentTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentTypesResponse struct{}"
	}

	return strings.Join([]string{"ListIncidentTypesResponse", string(data)}, " ")
}
