package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerlessPodInfo 实例基本信息
type ServerlessPodInfo struct {

	// 实例名称
	PodName *string `json:"pod_name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 所属集群
	ClusterName *string `json:"cluster_name,omitempty"`

	// CPU使用量
	Cpu *string `json:"cpu,omitempty"`

	// 内存使用量
	Memory *string `json:"memory,omitempty"`

	// cpu限制
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// 内存限制
	MemoryLimit *string `json:"memory_limit,omitempty"`

	// 实例 IP
	PodIp *string `json:"pod_ip,omitempty"`

	// 防护状态，包含如下2种。 - closed ：未防护。 - opened ：防护中。 - protection_exception ：防护异常。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// Serverless安全检测结果，包含如下4种。 - undetected ：未检测。 - clean ：无风险。 - risk ：有风险。 - scanning ：检测中。
	DetectResult *string `json:"detect_result,omitempty"`

	// Pod状态，包含以下几种 -Pending：pod 已被 Kubernetes 系统接受，但尚未创建一个或多个容器镜像 -Running：pod 已经绑定到一个节点，并且所有的容器都已经创建完毕 -Succeeded：pod 中的所有容器都已成功终止，不会重新启动 -Failed：Pod 中的所有容器都已终止，并且至少有一个容器因故障而终止 -Unknown：由于某种原因无法获取 pod 的状态，通常是由于与 pod 的主机通信时出错
	Status *interface{} `json:"status,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ServerlessPodInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerlessPodInfo struct{}"
	}

	return strings.Join([]string{"ServerlessPodInfo", string(data)}, " ")
}
