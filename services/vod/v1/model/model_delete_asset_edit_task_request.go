package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetEditTaskRequest Request Object
type DeleteAssetEditTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteAssetEditTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetEditTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetEditTaskRequest", string(data)}, " ")
}
