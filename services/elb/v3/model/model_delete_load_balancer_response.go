package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteLoadBalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerResponse", string(data)}, " ")
}
