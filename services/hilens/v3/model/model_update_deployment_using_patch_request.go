package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeploymentUsingPatchRequest Request Object
type UpdateDeploymentUsingPatchRequest struct {

	// 部署ID
	DeploymentId string `json:"deployment_id"`

	Body *DeploymentPatchRequest `json:"body,omitempty"`
}

func (o UpdateDeploymentUsingPatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentUsingPatchRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentUsingPatchRequest", string(data)}, " ")
}
