package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListItaSubJobsResponse Response Object
type ListItaSubJobsResponse struct {

	// 任务列表总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 任务列表。
	Jobs           *[]JobDetailInfo `json:"jobs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListItaSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListItaSubJobsResponse struct{}"
	}

	return strings.Join([]string{"ListItaSubJobsResponse", string(data)}, " ")
}
