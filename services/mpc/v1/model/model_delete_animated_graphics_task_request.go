package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAnimatedGraphicsTaskRequest Request Object
type DeleteAnimatedGraphicsTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteAnimatedGraphicsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteAnimatedGraphicsTaskRequest", string(data)}, " ")
}
