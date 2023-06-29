package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBatchAsyncJobsResponse Response Object
type UpdateBatchAsyncJobsResponse struct {

	// 批量更新指定ID异步任务响应体。
	Jobs           *[]AsyncUpdateJobResp `json:"jobs,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateBatchAsyncJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBatchAsyncJobsResponse struct{}"
	}

	return strings.Join([]string{"UpdateBatchAsyncJobsResponse", string(data)}, " ")
}
