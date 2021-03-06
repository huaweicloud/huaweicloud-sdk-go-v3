package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateIpGroupRequestBody struct {
	Ipgroup *UpdateIpGroupOption `json:"ipgroup"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}
