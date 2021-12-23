package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDeploymentGroupRequest struct {
	// 主机组ID

	GroupId string `json:"group_id"`

	Body *DeploymentGroupUpdateRequest `json:"body,omitempty"`
}

func (o UpdateDeploymentGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentGroupRequest", string(data)}, " ")
}
