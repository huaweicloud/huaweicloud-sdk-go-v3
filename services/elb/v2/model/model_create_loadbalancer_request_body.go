package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateLoadbalancerRequestBody struct {
	Loadbalancer *CreateLoadbalancerReq `json:"loadbalancer"`
}

func (o CreateLoadbalancerRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerRequestBody", string(data)}, " ")
}
