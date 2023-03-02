package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVideoCuttingTaskResponse struct {

	// 任务唯一标识
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoCuttingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoCuttingTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoCuttingTaskResponse", string(data)}, " ")
}
