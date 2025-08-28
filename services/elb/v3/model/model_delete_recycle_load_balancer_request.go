package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecycleLoadBalancerRequest Request Object
type DeleteRecycleLoadBalancerRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteRecycleLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecycleLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecycleLoadBalancerRequest", string(data)}, " ")
}
