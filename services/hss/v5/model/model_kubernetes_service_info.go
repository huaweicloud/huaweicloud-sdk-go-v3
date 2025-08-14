package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KubernetesServiceInfo 服务列表对象
type KubernetesServiceInfo struct {

	// id
	Id *string `json:"id,omitempty"`

	// 服务名称
	Name *string `json:"name,omitempty"`

	// 端点名称
	EndpointName *string `json:"endpoint_name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 创建时间戳
	CreationTimestamp *int64 `json:"creation_timestamp,omitempty"`

	// 服务类型（访问方式）
	Type *string `json:"type,omitempty"`

	// 集群IP
	ClusterIp *string `json:"cluster_ip,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群类型，包含以下几种： - k8s：原生集群 - cce：CCE集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o KubernetesServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KubernetesServiceInfo struct{}"
	}

	return strings.Join([]string{"KubernetesServiceInfo", string(data)}, " ")
}
