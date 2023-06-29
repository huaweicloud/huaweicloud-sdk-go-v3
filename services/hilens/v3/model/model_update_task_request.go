package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskRequest Request Object
type UpdateTaskRequest struct {

	// 部署ID，从专业版HiLens控制台部署管理[获取部署列表](getDeploymentListUsingGET.xml)获取
	DeploymentId string `json:"deployment_id"`

	// 作业ID，从专业版HiLens控制台作业管理[获取作业列表](listTasksUsingGET.xml)获取
	TaskId string `json:"task_id"`

	Body *TaskRequest `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
