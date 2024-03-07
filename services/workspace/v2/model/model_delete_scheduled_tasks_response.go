package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduledTasksResponse Response Object
type DeleteScheduledTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduledTasksResponse", string(data)}, " ")
}
