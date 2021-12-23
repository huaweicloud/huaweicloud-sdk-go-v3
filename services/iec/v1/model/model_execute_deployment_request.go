package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteDeploymentRequest struct {
	// 部署计划ID。  约束： - 该接口只能执行指定名称（name）创建的部署计划。

	DeploymentId string `json:"deployment_id"`
}

func (o ExecuteDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeploymentRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDeploymentRequest", string(data)}, " ")
}
