package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInfoVo struct {

	// **参数解释**： 分类。 **默认取值**： VacuumFull
	Category *string `json:"category,omitempty"`

	// **参数解释**： 描述信息。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 任务类型。 **默认取值**： Window：周期型任务； Date：单次型任务；
	Type *string `json:"type,omitempty"`

	// **参数解释**： 任务ID。 **默认取值**： 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**： 任务名称。 **默认取值**： 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 任务开始时间。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 任务结束时间。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 任务时间窗。 **默认取值**： 不涉及。
	WhiteList *[]DateInfo `json:"white_list,omitempty"`

	// **参数解释**： 任务模式。 **默认取值**： manual：指定目标； auto：自动；
	VacuumMode *string `json:"vacuum_mode,omitempty"`

	// **参数解释**： 自动Vacuum目标。 **默认取值**： user_vacuumfull：用户表VacuumFull； system_vacuum：系统表VacuumFull；
	VacuumTarget *string `json:"vacuum_target,omitempty"`

	// **参数解释**： 是否暂停。 **默认取值**： 0：否； 1：是；
	IsPaused *int32 `json:"is_paused,omitempty"`

	// **参数解释**： 膨胀率，单位为百分比。 **默认取值**： 不涉及。
	VacuumThreshold *string `json:"vacuum_threshold,omitempty"`

	// **参数解释**： 目标表可回收空间。 **默认取值**： 不涉及。
	VacuumRetrievingSpace *string `json:"vacuum_retrieving_space,omitempty"`

	// **参数解释**： 优先Vacuum目标。 **默认取值**： 不涉及。
	VacuumPriority *[]TableInfoOpen `json:"vacuum_priority,omitempty"`

	// **参数解释**： 时区信息。 **默认取值**： 一般为null。
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o TaskInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfoVo struct{}"
	}

	return strings.Join([]string{"TaskInfoVo", string(data)}, " ")
}
