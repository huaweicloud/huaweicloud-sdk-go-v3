package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerRequest Request Object
type DeleteLoadBalancerRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerRequest", string(data)}, " ")
}
