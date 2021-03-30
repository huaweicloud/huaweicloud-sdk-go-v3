package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerTagsRequest", string(data)}, " ")
}
