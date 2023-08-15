package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskResponse Response Object
type UpdateTaskResponse struct {

	// code
	Code *string `json:"code,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	TaskInfo       *TaskInfo `json:"taskInfo,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskResponse", string(data)}, " ")
}
