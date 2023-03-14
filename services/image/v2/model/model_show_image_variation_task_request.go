package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageVariationTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowImageVariationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageVariationTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowImageVariationTaskRequest", string(data)}, " ")
}
