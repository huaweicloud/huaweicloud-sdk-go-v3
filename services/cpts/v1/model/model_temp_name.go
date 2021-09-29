package model

import (
	"encoding/json"

	"strings"
)

type TempName struct {
	// name

	Name *string `json:"name,omitempty"`
}

func (o TempName) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TempName struct{}"
	}

	return strings.Join([]string{"TempName", string(data)}, " ")
}
