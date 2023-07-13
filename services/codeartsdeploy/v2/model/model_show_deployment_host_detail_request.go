package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeploymentHostDetailRequest Request Object
type ShowDeploymentHostDetailRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`
}

func (o ShowDeploymentHostDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentHostDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentHostDetailRequest", string(data)}, " ")
}
