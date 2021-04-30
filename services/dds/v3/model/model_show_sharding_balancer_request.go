package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowShardingBalancerRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowShardingBalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowShardingBalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowShardingBalancerRequest", string(data)}, " ")
}
