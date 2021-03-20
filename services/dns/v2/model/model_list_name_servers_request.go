package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNameServersRequest struct {
	Type *string `json:"type,omitempty"`

	Region *string `json:"region,omitempty"`
}

func (o ListNameServersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNameServersRequest struct{}"
	}

	return strings.Join([]string{"ListNameServersRequest", string(data)}, " ")
}
