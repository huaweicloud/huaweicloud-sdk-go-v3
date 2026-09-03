package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTasksNewResponse Response Object
type AddTasksNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddTasksNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTasksNewResponse struct{}"
	}

	return strings.Join([]string{"AddTasksNewResponse", string(data)}, " ")
}
