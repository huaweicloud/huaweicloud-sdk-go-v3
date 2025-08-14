package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowK8sPodDetailResponse Response Object
type ShowK8sPodDetailResponse struct {

	// pod名称
	PodName *string `json:"pod_name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 所属集群
	ClusterName *string `json:"cluster_name,omitempty"`

	// 所属节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 标签
	Label *string `json:"label,omitempty"`

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

	// Pod状态，包含以下几种 -Pending：pod已被Kubernetes系统接受，但尚未创建一个或多个容器镜像 -Running：pod已经绑定到一个节点，并且所有的容器都已经创建完毕 -Succeeded：pod中的所有容器都已成功终止，不会重新启动 -Failed：pod中的所有容器都已终止，并且至少有一个容器因故障而终止 -Unknown：由于某种原因无法获取pod的状态，通常是由于与pod的主机通信时出错
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// pod容器列表
	Containers     *[]ContainerBaseInfo `json:"containers,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowK8sPodDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowK8sPodDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowK8sPodDetailResponse", string(data)}, " ")
}
