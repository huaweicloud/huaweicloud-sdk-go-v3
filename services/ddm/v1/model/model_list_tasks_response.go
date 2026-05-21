package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponse Response Object
type ListTasksResponse struct {

	// 分页偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 一页数量。
	Limit *int32 `json:"limit,omitempty"`

	// 总条数。
	Total *int32 `json:"total,omitempty"`

	// 任务列表。
	Jobs           *[]JobItem `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
