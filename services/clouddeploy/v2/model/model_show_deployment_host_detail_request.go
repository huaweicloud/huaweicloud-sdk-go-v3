package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeploymentHostDetailRequest struct {
	// 主机组id

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
