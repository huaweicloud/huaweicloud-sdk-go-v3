package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentRequest Request Object
type DeleteDeploymentRequest struct {

	// 部署计划ID。
	DeploymentId string `json:"deployment_id"`
}

func (o DeleteDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentRequest", string(data)}, " ")
}
