package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponse Response Object
type ListTasksResponse struct {

	// 任务总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 任务名称列表。
	Actions *[]string `json:"actions,omitempty"`

	// 任务列表。
	Tasks          *[]Task `json:"tasks,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
