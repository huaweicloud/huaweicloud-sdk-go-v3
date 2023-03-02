package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoShotSplitTaskRequest struct {

	// 待查询任务ID
	TaskId string `json:"task_id"`
}

func (o ShowVideoShotSplitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoShotSplitTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoShotSplitTaskRequest", string(data)}, " ")
}
