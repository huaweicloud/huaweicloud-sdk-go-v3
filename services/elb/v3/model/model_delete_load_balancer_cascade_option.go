package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerCascadeOption **参数解释**：级联删除负载均衡器参数对象。  **约束限制**：不涉及
type DeleteLoadBalancerCascadeOption struct {

	// **参数解释**：是否删除关联的后端服务器组。  **约束限制**： - 共享型负载均衡器仅支持传参为true。 [- 若开启多挂特性，且后端服务器组关联了多个LB，则无论传入何值，都不会删除后端服务器组。](tag: hws,hws_hk)  **取值范围**： - true：删除该后端服务器组。 - false：仅解绑后端服务器组，不删除。  **默认取值**：true
	UnboundedPool *bool `json:"unbounded_pool,omitempty"`

	// **参数解释**：删除负载均衡器后是否删除关联的公网IP。  **约束限制**：不涉及          **取值范围**： - true：删除关联的EIP。 - false：仅解绑关联的EIP，不删除。  **默认取值**：false
	PublicIp *bool `json:"public_ip,omitempty"`
}

func (o DeleteLoadBalancerCascadeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerCascadeOption struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerCascadeOption", string(data)}, " ")
}
