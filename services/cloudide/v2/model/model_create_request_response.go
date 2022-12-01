package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRequestResponse struct {

	// the unique id of the request
	RequestId *string `json:"request_id,omitempty"`

	Status         *RequestStatus `json:"status,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestResponse struct{}"
	}

	return strings.Join([]string{"CreateRequestResponse", string(data)}, " ")
}
