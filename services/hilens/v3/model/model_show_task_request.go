package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTaskRequest struct {

	// 部署ID，从专业版HiLens控制台部署管理[获取部署列表](getDeploymentListUsingGET.xml)获取
	DeploymentId string `json:"deployment_id"`

	// 作业ID，从专业版HiLens控制台作业管理[获取作业列表](listTasksUsingGET.xml)获取
	TaskId string `json:"task_id"`
}

func (o ShowTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
