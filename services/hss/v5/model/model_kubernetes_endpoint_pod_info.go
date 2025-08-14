package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KubernetesEndpointPodInfo 端点关联Pod映射对象
type KubernetesEndpointPodInfo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 关联端点ID
	EndpointId *string `json:"endpoint_id,omitempty"`

	// pod IP
	PodIp *string `json:"pod_ip,omitempty"`

	// Pod名称
	PodName *string `json:"pod_name,omitempty"`

	// 是否可用
	Available *bool `json:"available,omitempty"`
}

func (o KubernetesEndpointPodInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KubernetesEndpointPodInfo struct{}"
	}

	return strings.Join([]string{"KubernetesEndpointPodInfo", string(data)}, " ")
}
