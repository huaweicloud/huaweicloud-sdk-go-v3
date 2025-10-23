package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponse Response Object
type ListTasksResponse struct {

	// 本次分页查询响应的任务条数
	Count *int64 `json:"count,omitempty"`

	// 下一页的marker
	NextMarker *string `json:"next_marker,omitempty"`

	// 任务列表
	Tasks          *[]TaskEntity `json:"tasks,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
