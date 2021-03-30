package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteLoadbalancerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsResponse", string(data)}, " ")
}
