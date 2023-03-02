package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVideoShotSplitTaskResponse struct {

	// 任务唯一标识
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoShotSplitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoShotSplitTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoShotSplitTaskResponse", string(data)}, " ")
}
