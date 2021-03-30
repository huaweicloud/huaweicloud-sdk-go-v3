package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadbalancerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerTagsResponse", string(data)}, " ")
}
