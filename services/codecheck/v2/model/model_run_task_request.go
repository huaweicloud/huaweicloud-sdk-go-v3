package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`

	Body *RunRequestV2 `json:"body,omitempty" xml:"body"`
}

func (o RunTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTaskRequest struct{}"
	}

	return strings.Join([]string{"RunTaskRequest", string(data)}, " ")
}
