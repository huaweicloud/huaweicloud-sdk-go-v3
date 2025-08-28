package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusHealthMonitor **参数解释**：LB状态树的后端服务器组健康检查器状态信息。
type LoadBalancerStatusHealthMonitor struct {

	// **参数解释**：健康检查器协议类型。  **取值范围**：TCP、UDP_CONNECT、HTTP。
	Type *string `json:"type,omitempty"`

	// **参数解释**：健康检查器ID。  **取值范围**：不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：健康检查器名称。  **取值范围**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：健康检查器的配置状态。  **取值范围**：ACTIVE表示使用中。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
}

func (o LoadBalancerStatusHealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusHealthMonitor struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusHealthMonitor", string(data)}, " ")
}
