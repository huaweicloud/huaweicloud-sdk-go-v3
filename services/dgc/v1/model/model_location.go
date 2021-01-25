package model

import (
	"encoding/json"

	"strings"
)

type Location struct {
	X *int32 `json:"x,omitempty"`
	Y *int32 `json:"y,omitempty"`
}

func (o Location) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}
