package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVideoObjectMaskingTaskResponse struct {

	// 任务唯一标识
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoObjectMaskingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoObjectMaskingTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoObjectMaskingTaskResponse", string(data)}, " ")
}
