package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataobjectRelationResponse Response Object
type ListDataobjectRelationResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// current page count
	Limit *int32 `json:"limit,omitempty"`

	// current page size
	Offset *int32 `json:"offset,omitempty"`

	// list of informations of dataobject
	Data *[]DataobjectInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDataobjectRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataobjectRelationResponse struct{}"
	}

	return strings.Join([]string{"ListDataobjectRelationResponse", string(data)}, " ")
}
