package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeploymentHostRequest struct {
	// 主机组id

	GroupId string `json:"group_id"`

	Body *DeploymentHost `json:"body,omitempty"`
}

func (o CreateDeploymentHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentHostRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentHostRequest", string(data)}, " ")
}
