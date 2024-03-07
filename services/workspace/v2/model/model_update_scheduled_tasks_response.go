package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledTasksResponse Response Object
type UpdateScheduledTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTasksResponse", string(data)}, " ")
}
