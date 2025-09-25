package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectImageContainersResponseInfo 受漏洞影响镜像的关联容器
type VulAffectImageContainersResponseInfo struct {

	// 容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 容器ID
	ContainerId *string `json:"container_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 节点名称
	HostName *string `json:"host_name,omitempty"`

	// 公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 私网IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// pod名称
	PodName *string `json:"pod_name,omitempty"`

	// pod id
	PodId *string `json:"pod_id,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o VulAffectImageContainersResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectImageContainersResponseInfo struct{}"
	}

	return strings.Join([]string{"VulAffectImageContainersResponseInfo", string(data)}, " ")
}
