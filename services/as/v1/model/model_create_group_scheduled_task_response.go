package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupScheduledTaskResponse Response Object
type CreateGroupScheduledTaskResponse struct {

	// 计划任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGroupScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupScheduledTaskResponse", string(data)}, " ")
}
