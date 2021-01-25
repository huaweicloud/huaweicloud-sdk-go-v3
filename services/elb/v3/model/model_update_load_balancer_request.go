package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateLoadBalancerRequest struct {
	LoadbalancerId string                         `json:"loadbalancer_id"`
	Body           *UpdateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o UpdateLoadBalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerRequest", string(data)}, " ")
}
