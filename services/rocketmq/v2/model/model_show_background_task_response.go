package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackgroundTaskResponse Response Object
type ShowBackgroundTaskResponse struct {

	// **参数解释**： 任务数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskCount *string `json:"task_count,omitempty"`

	// **参数解释**： 任务列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tasks          *[]ListBackgroundTasksRespTasks `json:"tasks,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowBackgroundTaskResponse", string(data)}, " ")
}
