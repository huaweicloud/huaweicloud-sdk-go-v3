package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeploymentGroupRequest struct {
	Body *DeploymentGroup `json:"body,omitempty" xml:"body"`
}

func (o CreateDeploymentGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentGroupRequest", string(data)}, " ")
}
