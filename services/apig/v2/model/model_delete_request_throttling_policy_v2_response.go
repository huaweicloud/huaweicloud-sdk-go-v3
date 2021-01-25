package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRequestThrottlingPolicyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRequestThrottlingPolicyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"DeleteRequestThrottlingPolicyV2Response", string(data)}, " ")
}
