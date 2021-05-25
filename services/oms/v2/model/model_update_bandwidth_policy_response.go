package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateBandwidthPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateBandwidthPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateBandwidthPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthPolicyResponse", string(data)}, " ")
}
