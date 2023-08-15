package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskRelatedTestCaseResponse Response Object
type UpdateTaskRelatedTestCaseResponse struct {

	// code
	Code *string `json:"code,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	TaskInfo       *TaskInfo `json:"taskInfo,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateTaskRelatedTestCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRelatedTestCaseResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskRelatedTestCaseResponse", string(data)}, " ")
}
