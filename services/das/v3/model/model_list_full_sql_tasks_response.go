package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlTasksResponse Response Object
type ListFullSqlTasksResponse struct {

	// SQL洞察任务列表。
	Tasks *[]FullSqlTask `json:"tasks,omitempty"`

	// 总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFullSqlTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlTasksResponse struct{}"
	}

	return strings.Join([]string{"ListFullSqlTasksResponse", string(data)}, " ")
}
