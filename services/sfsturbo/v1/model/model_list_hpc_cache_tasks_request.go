package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHpcCacheTasksRequest Request Object
type ListHpcCacheTasksRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 任务类型
	Type *string `json:"type,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// offset，默认值为 0
	Offset *int64 `json:"offset,omitempty"`

	// limit，默认值为 20
	Limit *int64 `json:"limit,omitempty"`
}

func (o ListHpcCacheTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHpcCacheTasksRequest struct{}"
	}

	return strings.Join([]string{"ListHpcCacheTasksRequest", string(data)}, " ")
}
