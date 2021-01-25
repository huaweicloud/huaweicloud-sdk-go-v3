package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateLoadBalancerRequestBody struct {
	Loadbalancer *UpdateLoadBalancerOption `json:"loadbalancer"`
}

func (o UpdateLoadBalancerRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerRequestBody", string(data)}, " ")
}
