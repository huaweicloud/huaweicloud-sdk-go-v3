package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteLoadBalancerRequest struct {
	// 负载均衡器ID。

	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadBalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerRequest", string(data)}, " ")
}
