package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SetBalancerWindowRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BalancerActiveWindow `json:"body,omitempty"`
}

func (o SetBalancerWindowRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetBalancerWindowRequest struct{}"
	}

	return strings.Join([]string{"SetBalancerWindowRequest", string(data)}, " ")
}
