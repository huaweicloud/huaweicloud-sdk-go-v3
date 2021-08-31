package model

import (
	"encoding/json"

	"strings"
)

// json
type ShowProcessResqJson struct {
	// details

	Details *[]ShowProcessResqJsonDetails `json:"details,omitempty"`
	// process_status

	ProcessStatus *int32 `json:"process_status,omitempty"`
}

func (o ShowProcessResqJson) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProcessResqJson struct{}"
	}

	return strings.Join([]string{"ShowProcessResqJson", string(data)}, " ")
}
