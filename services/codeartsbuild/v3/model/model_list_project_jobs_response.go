package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectJobsResponse Response Object
type ListProjectJobsResponse struct {

	// 任务列表
	Jobs *[]Job `json:"jobs,omitempty"`

	// 任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectJobsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectJobsResponse", string(data)}, " ")
}
