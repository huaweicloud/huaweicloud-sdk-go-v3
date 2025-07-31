package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KubernetesServicePortInfo Service关联端口列表对象
type KubernetesServicePortInfo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 关联服务 ID
	ServiceId *string `json:"service_id,omitempty"`

	// 端口名
	Name *string `json:"name,omitempty"`

	// 服务协议
	Protocol *string `json:"protocol,omitempty"`

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 后端容器端口
	TargetPort *string `json:"target_port,omitempty"`

	// 节点端口
	NodePort *int32 `json:"node_port,omitempty"`
}

func (o KubernetesServicePortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KubernetesServicePortInfo struct{}"
	}

	return strings.Join([]string{"KubernetesServicePortInfo", string(data)}, " ")
}
