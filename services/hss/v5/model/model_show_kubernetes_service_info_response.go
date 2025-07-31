package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKubernetesServiceInfoResponse Response Object
type ShowKubernetesServiceInfoResponse struct {

	// 服务ID
	Id *string `json:"id,omitempty"`

	// 服务名称
	Name *string `json:"name,omitempty"`

	// 端点名称
	EndpointName *string `json:"endpoint_name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 创建时间戳
	CreationTimestamp *int64 `json:"creation_timestamp,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 标签
	Labels *string `json:"labels,omitempty"`

	// 服务类型（访问方式）
	Type *string `json:"type,omitempty"`

	// 集群IP
	ClusterIp *string `json:"cluster_ip,omitempty"`

	// 选择器
	Selector *string `json:"selector,omitempty"`

	// 会话亲和性
	SessionAffinity *string `json:"session_affinity,omitempty"`

	// Service关联端口列表
	ServicePortList *[]KubernetesServicePortInfo `json:"service_port_list,omitempty"`
	HttpStatusCode  int                          `json:"-"`
}

func (o ShowKubernetesServiceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKubernetesServiceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowKubernetesServiceInfoResponse", string(data)}, " ")
}
