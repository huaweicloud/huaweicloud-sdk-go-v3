package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectProcessTaskResponse Response Object
type CreateObjectProcessTaskResponse struct {

	// 转码工作流任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateObjectProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateObjectProcessTaskResponse", string(data)}, " ")
}
