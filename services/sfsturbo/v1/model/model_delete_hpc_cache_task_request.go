package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHpcCacheTaskRequest Request Object
type DeleteHpcCacheTaskRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteHpcCacheTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHpcCacheTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteHpcCacheTaskRequest", string(data)}, " ")
}
