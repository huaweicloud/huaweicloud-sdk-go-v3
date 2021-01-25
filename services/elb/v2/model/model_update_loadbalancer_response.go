package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLoadbalancerResponse struct {
	Loadbalancer   *LoadbalancerResp `json:"loadbalancer,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateLoadbalancerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerResponse", string(data)}, " ")
}
