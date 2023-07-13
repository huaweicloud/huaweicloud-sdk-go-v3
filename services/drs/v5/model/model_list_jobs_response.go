package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsResponse Response Object
type ListJobsResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount int32 `json:"total_count"`

	// 任务信息列表。
	Jobs           []JobListResp `json:"jobs"`
	HttpStatusCode int           `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
