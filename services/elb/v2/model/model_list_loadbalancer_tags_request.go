package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListLoadbalancerTagsRequest struct {
}

func (o ListLoadbalancerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerTagsRequest", string(data)}, " ")
}
