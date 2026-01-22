package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksResponse Response Object
type ListScheduledTasksResponse struct {

	// **参数解释**： 任务总数。 **取值范围**： 不涉及。
	JobCount *string `json:"job_count,omitempty"`

	// **参数解释**： 任务列表。 **取值范围**： 不涉及。
	Jobs           *[]ScheduledTaskEntity `json:"jobs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksResponse", string(data)}, " ")
}
