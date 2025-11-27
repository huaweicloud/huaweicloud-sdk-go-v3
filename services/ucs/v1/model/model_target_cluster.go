package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TargetCluster struct {

	// 目标集群名称
	Name *string `json:"name,omitempty"`

	// 在该集群创建的副本数
	Replicas *string `json:"replicas,omitempty"`
}

func (o TargetCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetCluster struct{}"
	}

	return strings.Join([]string{"TargetCluster", string(data)}, " ")
}
