package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLoadbalancerRequest struct {
	// 负载均衡器ID

	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerRequest", string(data)}, " ")
}
