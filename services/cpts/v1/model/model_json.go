package model

import (
	"encoding/json"

	"strings"
)

// json
type Json struct {
	// details

	Details *[]JsonDetails `json:"details,omitempty"`
	// process_status

	ProcessStatus *int32 `json:"process_status,omitempty"`
}

func (o Json) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Json struct{}"
	}

	return strings.Join([]string{"Json", string(data)}, " ")
}
