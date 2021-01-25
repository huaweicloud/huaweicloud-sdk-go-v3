package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateLoadbalancerRequestBody struct {
	Loadbalancer *UpdateLoadbalancerReq `json:"loadbalancer"`
}

func (o UpdateLoadbalancerRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerRequestBody", string(data)}, " ")
}
