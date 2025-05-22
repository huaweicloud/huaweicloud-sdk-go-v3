package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecycleLoadBalancerRequest Request Object
type DeleteRecycleLoadBalancerRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteRecycleLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecycleLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecycleLoadBalancerRequest", string(data)}, " ")
}
