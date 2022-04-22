package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AomMappingRuleInfo struct {

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 工作负载
	Deployments []string `json:"deployments"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 接入规则详情
	Files []AomMappingfilesInfo `json:"files"`
}

func (o AomMappingRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AomMappingRuleInfo struct{}"
	}

	return strings.Join([]string{"AomMappingRuleInfo", string(data)}, " ")
}
