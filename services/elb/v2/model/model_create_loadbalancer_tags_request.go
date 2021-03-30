package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *CreateLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o CreateLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsRequest", string(data)}, " ")
}
