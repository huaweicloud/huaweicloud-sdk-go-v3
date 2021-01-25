package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateLoadBalancerRequest struct {
	Body *CreateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadBalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerRequest", string(data)}, " ")
}
