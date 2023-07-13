package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAsyncJobsResponse Response Object
type ListAsyncJobsResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount int32 `json:"total_count"`

	// 所有批量异步创建任务响应体。
	Jobs           []AsyncJobResp `json:"jobs"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAsyncJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncJobsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncJobsResponse", string(data)}, " ")
}
