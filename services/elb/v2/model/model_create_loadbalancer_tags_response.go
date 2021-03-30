package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateLoadbalancerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsResponse", string(data)}, " ")
}
