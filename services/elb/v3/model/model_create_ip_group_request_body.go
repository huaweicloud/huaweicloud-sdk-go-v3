package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateIpGroupRequestBody struct {
	Ipgroup *CreateIpGroupOption `json:"ipgroup"`
}

func (o CreateIpGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequestBody", string(data)}, " ")
}
