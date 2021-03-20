package model

import (
	"encoding/json"

	"strings"
)

type StacksConfig struct {
	Attributes *StacksAttribute `json:"attributes,omitempty"`

	Recipe *Recipe `json:"recipe,omitempty"`
}

func (o StacksConfig) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StacksConfig struct{}"
	}

	return strings.Join([]string{"StacksConfig", string(data)}, " ")
}
