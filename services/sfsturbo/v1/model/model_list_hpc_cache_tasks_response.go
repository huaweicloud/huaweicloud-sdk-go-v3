package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHpcCacheTasksResponse Response Object
type ListHpcCacheTasksResponse struct {

	// 任务详情
	Tasks *[]OneHpcCacheTaskInfoResp `json:"tasks,omitempty"`

	// 任务数量
	Count *int64 `json:"count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHpcCacheTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHpcCacheTasksResponse struct{}"
	}

	return strings.Join([]string{"ListHpcCacheTasksResponse", string(data)}, " ")
}
