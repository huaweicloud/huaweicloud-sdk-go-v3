package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailRequest Request Object
type ShowTaskDetailRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailRequest", string(data)}, " ")
}
