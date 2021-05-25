package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTopicAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicAccessPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyResponse", string(data)}, " ")
}
