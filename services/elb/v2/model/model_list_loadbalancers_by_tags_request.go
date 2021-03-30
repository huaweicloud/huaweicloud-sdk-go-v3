package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListLoadbalancersByTagsRequest struct {
	Body *ListLoadbalancersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListLoadbalancersByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLoadbalancersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersByTagsRequest", string(data)}, " ")
}
