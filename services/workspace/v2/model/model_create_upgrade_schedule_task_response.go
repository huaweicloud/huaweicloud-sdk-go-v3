package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUpgradeScheduleTaskResponse Response Object
type CreateUpgradeScheduleTaskResponse struct {

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUpgradeScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpgradeScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateUpgradeScheduleTaskResponse", string(data)}, " ")
}
