package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteLoadbalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadbalancerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerResponse", string(data)}, " ")
}
