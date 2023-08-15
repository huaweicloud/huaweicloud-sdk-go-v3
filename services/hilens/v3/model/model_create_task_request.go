package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequest Request Object
type CreateTaskRequest struct {

	// 部署ID，从专业版HiLens控制台部署管理[获取部署列表](getDeploymentListUsingGET.xml)获取
	DeploymentId string `json:"deployment_id"`

	Body *TaskRequest `json:"body,omitempty"`
}

func (o CreateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskRequest", string(data)}, " ")
}
