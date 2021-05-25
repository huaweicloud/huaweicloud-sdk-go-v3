package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLoadBalancerStatusRequest struct {
	// 所属负载均衡器id

	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadBalancerStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerStatusRequest", string(data)}, " ")
}
