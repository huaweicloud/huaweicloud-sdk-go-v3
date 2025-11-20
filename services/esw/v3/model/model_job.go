package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Job struct {

	// - 参数解释：任务的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Id string `json:"id"`

	// - 参数解释：当前任务的名称。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name string `json:"name"`

	// - 参数解释：任务状态。 - 约束限制：不涉及。 - 取值范围：   - RUNNING：运行中   - FAILED：失败   - COMPLETED：已完成 - 默认取值：不涉及。
	Status string `json:"status"`

	// - 参数解释：任务开始时间。 - 约束限制：UTC时间，格式为yyyy-MM-ddTHH:mm:ss。 - 取值范围：不涉及。 - 默认取值：不涉及。
	BeginTime string `json:"begin_time"`

	// - 参数解释：任务结束时间。 - 约束限制：   - UTC时间，格式为yyyy-MM-ddTHH:mm:ss。   - 仅在任务状态为FAILED或者COMPLETED时可见 - 取值范围：不涉及。 - 默认取值：不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// - 参数解释：任务当前进度，以百分比展示。 - 约束限制：仅在任务状态为RUNNING时可见。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Process *string `json:"process,omitempty"`

	// - 参数解释：任务的失败原因。 - 约束限制：仅在任务状态为FAILED时显示。 - 取值范围：不涉及。 - 默认取值：不涉及。
	FailReason *string `json:"fail_reason,omitempty"`

	// - 参数解释：任务当前关联的资源ID。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ResourceId string `json:"resource_id"`

	// - 参数解释：任务当前关联的资源名称。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ResourceName string `json:"resource_name"`

	// - 参数解释：任务当前关联的资源类型。 - 约束限制：不涉及。 - 取值范围：   - instance：ESW实例 - 默认取值：不涉及。
	ResourceType string `json:"resource_type"`

	// - 参数解释：ESW实例所属项目ID。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ProjectId string `json:"project_id"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
