package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTaskActionRequest Request Object
type ExecuteTaskActionRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 任务id
	TaskId string `json:"task_id"`

	// 启动、调度、停止操作标识
	Action string `json:"action"`
}

func (o ExecuteTaskActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTaskActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTaskActionRequest", string(data)}, " ")
}
