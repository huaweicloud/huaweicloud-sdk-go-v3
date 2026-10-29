package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskCmetricsResponse Response Object
type ShowTaskCmetricsResponse struct {

	// **参数解释：** 旧版任务ID。即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。 **取值范围：** 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 任务名称。 **取值范围：** 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释：** 任务创建者用户ID。 **取值范围：** 不涉及。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释：** 代码仓地址。 **取值范围：** 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释：** 代码仓分支。 **取值范围：** 不涉及。
	GitBranch *string `json:"git_branch,omitempty"`

	// **参数解释：** 上一次检查时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastCheckTime *string `json:"last_check_time,omitempty"`

	// **参数解释：** 上一次执行时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastExecTime *string `json:"last_exec_time,omitempty"`

	// **参数解释：** 上次执行结束时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastCheckFinishTime *string `json:"last_check_finish_time,omitempty"`

	// **参数解释：** 废弃字段。 **约束限制：** 不涉及。 **取值范围：** 不涉及。
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释：** 任务创建时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	MetricInfo     *MetricInfo `json:"metric_info,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTaskCmetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskCmetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskCmetricsResponse", string(data)}, " ")
}
