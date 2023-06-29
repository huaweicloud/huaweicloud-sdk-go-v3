package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerForceRequest Request Object
type DeleteLoadBalancerForceRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadBalancerForceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerForceRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerForceRequest", string(data)}, " ")
}
