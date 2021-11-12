package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateLoadBalancerRequest struct {
	// 负载均衡器ID。

	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o UpdateLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerRequest", string(data)}, " ")
}
