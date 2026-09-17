package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryNodeResp 查询节点详情响应结构体
type QueryNodeResp struct {

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点ip
	InternalIp *string `json:"internal_ip,omitempty"`

	// 主机名
	Hostname *string `json:"hostname,omitempty"`

	Allocatable *NodeResourceDto `json:"allocatable,omitempty"`

	Capacity *NodeResourceDto `json:"capacity,omitempty"`

	AllocatedResources *NodeAllocatedResourceDto `json:"allocated_resources,omitempty"`

	// 状态，Ready or NotReady
	Status *string `json:"status,omitempty"`

	// 架构，amd64 or arm64
	Architecture *string `json:"architecture,omitempty"`

	// map类型，key为string,value为string
	Labels map[string]string `json:"labels,omitempty"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 内核版本
	KernelVersion *string `json:"kernel_version,omitempty"`

	// 操作系统版本
	OsImage *string `json:"os_image,omitempty"`

	// 容器运行时版本
	ContainerRuntimeVersion *string `json:"container_runtime_version,omitempty"`

	// k8s版本
	KubernetesVersion *string `json:"kubernetes_version,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o QueryNodeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryNodeResp struct{}"
	}

	return strings.Join([]string{"QueryNodeResp", string(data)}, " ")
}
