package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsResponse Response Object
type ListJobsResponse struct {

	// 任务列表
	Jobs *[]MainJob `json:"jobs,omitempty"`

	// 请求ID。 注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
