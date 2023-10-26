package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHpcCacheTaskRequest Request Object
type CreateHpcCacheTaskRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *CreateHpcCacheTaskReq `json:"body,omitempty"`
}

func (o CreateHpcCacheTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHpcCacheTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateHpcCacheTaskRequest", string(data)}, " ")
}
