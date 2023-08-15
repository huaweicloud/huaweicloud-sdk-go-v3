package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResolveTasksResponse Response Object
type ListResolveTasksResponse struct {

	// 解析任务列表。
	ResolveTasks *[]ListResolveTaskResult `json:"resolve_tasks,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListResolveTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolveTasksResponse struct{}"
	}

	return strings.Join([]string{"ListResolveTasksResponse", string(data)}, " ")
}
