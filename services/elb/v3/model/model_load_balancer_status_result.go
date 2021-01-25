package model

import (
	"encoding/json"

	"strings"
)

//
type LoadBalancerStatusResult struct {
	Loadbalancer *LoadBalancerStatus `json:"loadbalancer"`
}

func (o LoadBalancerStatusResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusResult struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusResult", string(data)}, " ")
}
