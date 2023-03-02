package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoCoverAnalysisTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowVideoCoverAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoCoverAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoCoverAnalysisTaskRequest", string(data)}, " ")
}
