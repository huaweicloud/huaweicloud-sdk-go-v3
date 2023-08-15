package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentGroupRequest Request Object
type CreateDeploymentGroupRequest struct {
	Body *DeploymentGroup `json:"body,omitempty"`
}

func (o CreateDeploymentGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentGroupRequest", string(data)}, " ")
}
