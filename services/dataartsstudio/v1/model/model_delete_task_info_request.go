package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskInfoRequest Request Object
type DeleteTaskInfoRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 任务id
	TaskId string `json:"task_id"`
}

func (o DeleteTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteTaskInfoRequest", string(data)}, " ")
}
