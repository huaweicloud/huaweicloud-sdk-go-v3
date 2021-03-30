package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchDeleteLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsRequest", string(data)}, " ")
}
