package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentHostRequest Request Object
type CreateDeploymentHostRequest struct {

	// 主机集群id
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
