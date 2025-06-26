package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceWebhookJobsResponse Response Object
type ListInstanceWebhookJobsResponse struct {

	// 触发器执行任务列表
	Jobs *[]Job `json:"jobs,omitempty"`

	// 触发器执行任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceWebhookJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceWebhookJobsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceWebhookJobsResponse", string(data)}, " ")
}
