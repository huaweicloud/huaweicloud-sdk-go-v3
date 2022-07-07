package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LB状态树的后端服务器组健康检查器状态信息。
type LoadBalancerStatusHealthMonitor struct {

	// 协议类型。取值：TCP、UDP_CONNECT或HTTP。
	Type *string `json:"type,omitempty"`

	// 健康检查器ID。
	Id *string `json:"id,omitempty"`

	// 健康检查器名称。
	Name *string `json:"name,omitempty"`

	// 健康检查器的配置状态。取值：ACTIVE表示使用中。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
}

func (o LoadBalancerStatusHealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusHealthMonitor struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusHealthMonitor", string(data)}, " ")
}
