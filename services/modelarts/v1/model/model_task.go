package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Task 训练作业的任务列表。
type Task struct {

	// **参数解释**：任务角色，该功能暂未支持。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Role *string `json:"role,omitempty"`

	Algorithm *TaskAlgorithm `json:"algorithm,omitempty"`

	TaskResource *TaskTaskResource `json:"task_resource,omitempty"`

	LogExportPath *TaskLogExportPath `json:"log_export_path,omitempty"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
