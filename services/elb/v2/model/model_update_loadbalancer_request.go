package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateLoadbalancerRequest struct {
	// 待更新的负载均衡器id

	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o UpdateLoadbalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerRequest", string(data)}, " ")
}
