package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateLoadbalancerResponse struct {
	Loadbalancer   *LoadbalancerResp `json:"loadbalancer,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateLoadbalancerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerResponse", string(data)}, " ")
}
