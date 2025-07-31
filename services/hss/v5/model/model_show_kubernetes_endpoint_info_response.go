package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKubernetesEndpointInfoResponse Response Object
type ShowKubernetesEndpointInfoResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// 端点名称
	Name *string `json:"name,omitempty"`

	// 服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 创建时间戳
	CreationTimestamp *int64 `json:"creation_timestamp,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 标签
	Labels *string `json:"labels,omitempty"`

	// 是否关联服务
	AssociationService *bool `json:"association_service,omitempty"`

	// 端点关联Pod映射
	EndpointPodList *[]KubernetesEndpointPodInfo `json:"endpoint_pod_list,omitempty"`

	// 端点关联端口列表
	EndpointPortList *[]KubernetesEndpointPortInfo `json:"endpoint_port_list,omitempty"`
	HttpStatusCode   int                           `json:"-"`
}

func (o ShowKubernetesEndpointInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKubernetesEndpointInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowKubernetesEndpointInfoResponse", string(data)}, " ")
}
