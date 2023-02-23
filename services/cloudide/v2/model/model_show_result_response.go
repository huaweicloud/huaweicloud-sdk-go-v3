package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResultResponse struct {
	RequestStatus *RequestStatus `json:"request_status,omitempty"`

	// task list
	Tasks *[]TaskModel `json:"tasks,omitempty"`

	// the type of the request
	RequestType    *string `json:"request_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResultResponse struct{}"
	}

	return strings.Join([]string{"ShowResultResponse", string(data)}, " ")
}
