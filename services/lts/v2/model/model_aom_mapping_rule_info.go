package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AomMappingRuleInfo struct {

	// 集群id
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name" xml:"cluster_name"`

	// 日志流前缀
	DeploymentsPrefix *string `json:"deployments_prefix,omitempty" xml:"deployments_prefix"`

	// 工作负载
	Deployments []string `json:"deployments" xml:"deployments"`

	// 命名空间
	Namespace string `json:"namespace" xml:"namespace"`

	// 容器名称
	ContainerName *string `json:"container_name,omitempty" xml:"container_name"`

	// 接入规则详情
	Files []AomMappingfilesInfo `json:"files" xml:"files"`
}

func (o AomMappingRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AomMappingRuleInfo struct{}"
	}

	return strings.Join([]string{"AomMappingRuleInfo", string(data)}, " ")
}
