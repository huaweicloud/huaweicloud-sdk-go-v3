package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentHostRequest Request Object
type DeleteDeploymentHostRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`
}

func (o DeleteDeploymentHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentHostRequest", string(data)}, " ")
}
