package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodBaseInfo pod基本信息
type PodBaseInfo struct {

	// pod名称
	PodName *string `json:"pod_name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 所属集群
	ClusterName *string `json:"cluster_name,omitempty"`

	// 所属节点名称
	NodeName *string `json:"node_name,omitempty"`

	// CPU使用量
	Cpu *string `json:"cpu,omitempty"`

	// 内存使用量
	Memory *string `json:"memory,omitempty"`

	// cpu限制
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// 内存限制
	MemoryLimit *string `json:"memory_limit,omitempty"`

	// 所属节点IP
	NodeIp *string `json:"node_ip,omitempty"`

	// Pod IP
	PodIp *string `json:"pod_ip,omitempty"`

	// Pod状态，包含以下几种 -Pending：pod 已被 Kubernetes 系统接受，但尚未创建一个或多个容器镜像 -Running：pod 已经绑定到一个节点，并且所有的容器都已经创建完毕 -Succeeded：pod 中的所有容器都已成功终止，不会重新启动 -Failed：Pod 中的所有容器都已终止，并且至少有一个容器因故障而终止 -Unknown：由于某种原因无法获取 pod 的状态，通常是由于与 pod 的主机通信时出错
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// regionId
	RegionId *string `json:"region_id,omitempty"`

	// ID
	Id *string `json:"id,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含以下几种： -k8s 原生集群 -cce CCE集群 -ali 阿里云集群 -tencent 腾讯云集群 -azure 微软云集群 -aws 亚马逊集群 -self_built_hw 华为云自建集群 -self_built_idc IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o PodBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodBaseInfo struct{}"
	}

	return strings.Join([]string{"PodBaseInfo", string(data)}, " ")
}
