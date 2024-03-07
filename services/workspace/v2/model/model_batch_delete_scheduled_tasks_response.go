package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScheduledTasksResponse Response Object
type BatchDeleteScheduledTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteScheduledTasksResponse", string(data)}, " ")
}
