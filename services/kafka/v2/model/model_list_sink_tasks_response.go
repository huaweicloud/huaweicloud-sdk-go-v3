package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSinkTasksResponse struct {

	// 转储任务列表。
	Tasks *[]ListSinkTasksRespTasks `json:"tasks,omitempty" xml:"tasks"`

	// 转储任务总数。
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number"`

	// 总的支持任务个数。
	MaxTasks *int32 `json:"max_tasks,omitempty" xml:"max_tasks"`

	// 任务总数的配额。
	QuotaTasks     *int32 `json:"quota_tasks,omitempty" xml:"quota_tasks"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSinkTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSinkTasksResponse struct{}"
	}

	return strings.Join([]string{"ListSinkTasksResponse", string(data)}, " ")
}
