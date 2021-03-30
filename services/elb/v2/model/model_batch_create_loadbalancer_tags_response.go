package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateLoadbalancerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsResponse", string(data)}, " ")
}
