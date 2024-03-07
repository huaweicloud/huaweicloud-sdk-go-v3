package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduledTasksReq 删除定时任务请求体。
type DeleteScheduledTasksReq struct {

	// 待删除的任务ID列表。
	TaskIds *[]string `json:"task_ids,omitempty"`
}

func (o DeleteScheduledTasksReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduledTasksReq struct{}"
	}

	return strings.Join([]string{"DeleteScheduledTasksReq", string(data)}, " ")
}
