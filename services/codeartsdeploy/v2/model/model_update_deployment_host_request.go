package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeploymentHostRequest Request Object
type UpdateDeploymentHostRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`

	Body *DeploymentHostRequest `json:"body,omitempty"`
}

func (o UpdateDeploymentHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentHostRequest", string(data)}, " ")
}
