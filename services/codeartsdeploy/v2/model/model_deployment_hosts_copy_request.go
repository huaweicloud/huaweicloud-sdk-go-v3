package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentHostsCopyRequest struct {

	// 主机id列表
	HostUuids []string `json:"host_uuids"`

	// 目标主机集群id
	TargetGroupId string `json:"target_group_id"`
}

func (o DeploymentHostsCopyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHostsCopyRequest struct{}"
	}

	return strings.Join([]string{"DeploymentHostsCopyRequest", string(data)}, " ")
}
