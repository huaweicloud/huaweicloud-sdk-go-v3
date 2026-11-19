package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelColdTableWaitingTasksResponse Response Object
type CancelColdTableWaitingTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelColdTableWaitingTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelColdTableWaitingTasksResponse struct{}"
	}

	return strings.Join([]string{"CancelColdTableWaitingTasksResponse", string(data)}, " ")
}
