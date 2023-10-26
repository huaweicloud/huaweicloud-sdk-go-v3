package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHpcCacheTaskResponse Response Object
type CreateHpcCacheTaskResponse struct {

	// 任务任务id
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateHpcCacheTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHpcCacheTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateHpcCacheTaskResponse", string(data)}, " ")
}
