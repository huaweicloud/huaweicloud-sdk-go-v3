package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataobjectRelationResponse Response Object
type DeleteDataobjectRelationResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// Error message
	RequestId *string `json:"request_id,omitempty"`

	// Error message
	Success *bool `json:"success,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// current page count
	Limit *int32 `json:"limit,omitempty"`

	// current page size
	Offset *int32 `json:"offset,omitempty"`

	Data *DataResponse `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataobjectRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataobjectRelationResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataobjectRelationResponse", string(data)}, " ")
}
