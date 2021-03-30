package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchCreateLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsRequest", string(data)}, " ")
}
