package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskResponse 训练作业的任务列表。
type TaskResponse struct {

	// 任务角色，该功能暂未支持。
	Role *string `json:"role,omitempty"`

	Algorithm *TaskResponseAlgorithm `json:"algorithm,omitempty"`

	TaskResource *FlavorResponse `json:"task_resource,omitempty"`

	LogExportPath *TaskResponseLogExportPath `json:"log_export_path,omitempty"`
}

func (o TaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskResponse struct{}"
	}

	return strings.Join([]string{"TaskResponse", string(data)}, " ")
}
