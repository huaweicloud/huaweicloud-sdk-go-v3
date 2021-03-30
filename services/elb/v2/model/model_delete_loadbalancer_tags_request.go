package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Key string `json:"key"`
}

func (o DeleteLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerTagsRequest", string(data)}, " ")
}
