package model

import (
	"encoding/json"

	"strings"
)

type ResourceDto struct {
	Limits *ResourceConfigDto `json:"limits,omitempty"`

	Requests *ResourceConfigDto `json:"requests,omitempty"`
}

func (o ResourceDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
