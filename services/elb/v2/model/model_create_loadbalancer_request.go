package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateLoadbalancerRequest struct {
	Body *CreateLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadbalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerRequest", string(data)}, " ")
}
