package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateRequestThrottlingPolicyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateRequestThrottlingPolicyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateRequestThrottlingPolicyV2Response", string(data)}, " ")
}
