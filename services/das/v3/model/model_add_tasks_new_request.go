package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTasksNewRequest Request Object
type AddTasksNewRequest struct {
	Body *AddTasksNewRequestBody `json:"body,omitempty"`
}

func (o AddTasksNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTasksNewRequest struct{}"
	}

	return strings.Join([]string{"AddTasksNewRequest", string(data)}, " ")
}
