package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImmediateJobsResponse Response Object
type ListImmediateJobsResponse struct {

	// 任务详情。
	Jobs *[]TaskDetailInfo `json:"jobs,omitempty"`

	// 任务总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  支持筛选的任务名称。  **取值范围**：  不涉及。
	ActionNames    *[]string `json:"action_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListImmediateJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImmediateJobsResponse struct{}"
	}

	return strings.Join([]string{"ListImmediateJobsResponse", string(data)}, " ")
}
