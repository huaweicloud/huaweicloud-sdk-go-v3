package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRefreshTasksResponse Response Object
type CreateRefreshTasksResponse struct {

	// 任务ID
	RefreshTask    *string `json:"refresh_task,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRefreshTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRefreshTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateRefreshTasksResponse", string(data)}, " ")
}
