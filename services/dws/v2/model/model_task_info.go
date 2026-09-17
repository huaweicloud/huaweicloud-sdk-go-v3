package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInfo struct {

	// **参数解释**： 任务名称。 **约束限制**： 不涉及。 **取值范围**： 非null。 **默认取值**： 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 描述信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 开始时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **约束限制**： 不涉及。 **取值范围**： 非null。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 执行计划。 **约束限制**： 不涉及。 **取值范围**： 非null。 **默认取值**： 不涉及。
	WhiteList *[]DateInfo `json:"white_list,omitempty"`

	// **参数解释**： 任务类型。 **约束限制**： 不涉及。 **取值范围**： Date：单次型任务； Window：周期型任务； **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 任务模式。 **约束限制**： 不涉及。 **取值范围**： 非null。 **默认取值**： 不涉及。
	VacuumMode *string `json:"vacuum_mode,omitempty"`

	// **参数解释**： 自动Vacuum目标。 **约束限制**： 不涉及。 **取值范围**： user_vacuumfull：用户表VacuumFull； system_vacuum：系统表VacuumFull； **默认取值**： 不涉及。
	VacuumTarget *string `json:"vacuum_target,omitempty"`

	// **参数解释**： 膨胀率，单位为百分比。 **约束限制**： 不涉及。 **取值范围**： 建议设置为当前集群空间使用率+10%，且最大不超过80%。 **默认取值**： 不涉及。
	VacuumThreshold *string `json:"vacuum_threshold,omitempty"`

	// **参数解释**： 目标表可回收空间。 **约束限制**： 不涉及。 **取值范围**： 非null。 **默认取值**： 不涉及。
	VacuumRetrievingSpace *string `json:"vacuum_retrieving_space,omitempty"`

	// **参数解释**： 优先级。 **约束限制**： 不涉及。 **取值范围**： 非null。 **默认取值**： 不涉及。
	VacuumPriority *string `json:"vacuum_priority,omitempty"`

	// **参数解释**： 时区偏移信息。 **约束限制**： 不涉及。 **取值范围**： -2659~+2459 **默认取值**： 不涉及。
	TimeZone *string `json:"time_zone,omitempty"`

	// **参数解释**： 优先级表信息。 **默认取值**： 不涉及。
	Priority *[]TableInfoOpen `json:"priority,omitempty"`

	// **参数解释**： 分类信息。 **默认取值**： Vacuum
	Category *string `json:"category,omitempty"`
}

func (o TaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}
