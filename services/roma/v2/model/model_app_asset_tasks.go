package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppAssetTasks struct {
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
}

func (o AppAssetTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAssetTasks struct{}"
	}

	return strings.Join([]string{"AppAssetTasks", string(data)}, " ")
}
