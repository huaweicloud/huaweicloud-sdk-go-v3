package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceJobsResponse Response Object
type ListInstanceJobsResponse struct {

	// 任务列表
	Jobs *[]JobDetail `json:"jobs,omitempty"`

	// 任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceJobsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceJobsResponse", string(data)}, " ")
}
