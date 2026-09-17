package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskStatusOpenResp struct {

	// **参数解释**： 任务ID。 **默认取值**： 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**： 任务截止时间。 **默认取值**： 不涉及。
	TaskDeadLine *string `json:"task_dead_line,omitempty"`

	// **参数解释**： 分类信息。 **默认取值**： VacuumFull
	Category *string `json:"category,omitempty"`

	// **参数解释**： 状态。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 剩余时间。 **默认取值**： 不涉及。
	TimeLeft *string `json:"time_left,omitempty"`

	// **参数解释**： 开始时间。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 完成百分比。 **默认取值**： 不涉及。
	FinishedPercentage *string `json:"finished_percentage,omitempty"`

	// **参数解释**： 任务ID。 **默认取值**： 不涉及。
	VacuumedSpace *string `json:"vacuumed_space,omitempty"`

	TableVacuumInfo *TableVacuumInfoOpen `json:"table_vacuum_info,omitempty"`

	// **参数解释**： Vacuum信息。 **默认取值**： 不涉及。
	VacuumInfo *[]interface{} `json:"vacuum_info,omitempty"`

	TableVacuumNumInfo *TableVacuumNumInfo `json:"table_vacuum_num_info,omitempty"`
}

func (o TaskStatusOpenResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatusOpenResp struct{}"
	}

	return strings.Join([]string{"TaskStatusOpenResp", string(data)}, " ")
}
