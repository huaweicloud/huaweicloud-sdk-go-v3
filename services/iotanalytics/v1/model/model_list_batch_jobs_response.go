package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBatchJobsResponse Response Object
type ListBatchJobsResponse struct {

	// 定时作业总个数。
	Count *int64 `json:"count,omitempty"`

	Jobs           *[]Job `json:"jobs,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBatchJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchJobsResponse struct{}"
	}

	return strings.Join([]string{"ListBatchJobsResponse", string(data)}, " ")
}
