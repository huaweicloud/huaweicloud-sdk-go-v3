package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetOperateRequestTasks struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`
}

func (o AssetOperateRequestTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetOperateRequestTasks struct{}"
	}

	return strings.Join([]string{"AssetOperateRequestTasks", string(data)}, " ")
}
