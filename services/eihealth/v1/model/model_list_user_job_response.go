package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserJobResponse Response Object
type ListUserJobResponse struct {

	// 作业列表。
	Jobs *[]JobListDto `json:"jobs,omitempty"`

	// 作业总数。
	Count *int32 `json:"count,omitempty"`

	// 运行中作业总数。
	RunningCount *int32 `json:"running_count,omitempty"`

	// 等待中作业总数。
	WaitingCount   *int32 `json:"waiting_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserJobResponse struct{}"
	}

	return strings.Join([]string{"ListUserJobResponse", string(data)}, " ")
}
