package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHpcCacheTaskRequest Request Object
type ShowHpcCacheTaskRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowHpcCacheTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHpcCacheTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowHpcCacheTaskRequest", string(data)}, " ")
}
