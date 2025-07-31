package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KubernetesEndpointPortInfo 端点关联端口列表对象
type KubernetesEndpointPortInfo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 关联端点 ID
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 端口名
	Name *string `json:"name,omitempty"`

	// 服务协议
	Protocol *string `json:"protocol,omitempty"`

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 应用协议
	AppProtocol *string `json:"app_protocol,omitempty"`
}

func (o KubernetesEndpointPortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KubernetesEndpointPortInfo struct{}"
	}

	return strings.Join([]string{"KubernetesEndpointPortInfo", string(data)}, " ")
}
