package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateClusterRequestDto struct {

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 集群描述
	Description *string `json:"description,omitempty"`

	ClusterNodeConfig *ClusterNodeConfig `json:"cluster_node_config,omitempty"`
}

func (o CreateClusterRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequestDto struct{}"
	}

	return strings.Join([]string{"CreateClusterRequestDto", string(data)}, " ")
}
