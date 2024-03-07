package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduledTasksResponse Response Object
type CreateScheduledTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduledTasksResponse", string(data)}, " ")
}
