package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLoadbalancerTagsRequest struct {
	// 负载均衡器所在的项目id

	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerTagsRequest", string(data)}, " ")
}
