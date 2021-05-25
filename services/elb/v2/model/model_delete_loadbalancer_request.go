package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteLoadbalancerRequest struct {
	// 负载均衡器id

	LoadbalancerId string `json:"loadbalancer_id"`
	// （不再支持）级联删除负载均衡器

	Cascade *bool `json:"cascade,omitempty"`
}

func (o DeleteLoadbalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerRequest", string(data)}, " ")
}
