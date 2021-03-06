package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateL7policiesRequestBody struct {
	L7policy *UpdateL7policyReq `json:"l7policy"`
}

func (o UpdateL7policiesRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7policiesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesRequestBody", string(data)}, " ")
}
