package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInfoRequest Request Object
type ShowTaskInfoRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskInfoRequest", string(data)}, " ")
}
