package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRequestResponse Response Object
type CreateRequestResponse struct {

	// the unique id of the request
	RequestId *string `json:"request_id,omitempty"`

	Status *RequestStatus `json:"status,omitempty"`

	// the number of tasks dispatched successfully
	DispatchedTaskNumber *int32 `json:"dispatched_task_number,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o CreateRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestResponse struct{}"
	}

	return strings.Join([]string{"CreateRequestResponse", string(data)}, " ")
}
