package model

import (
	"encoding/json"

	"strings"
)

type ShowProcessResqJsonDetails struct {
	// id

	Id *int32 `json:"id,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// status

	Status *int32 `json:"status,omitempty"`
	// cause

	Cause *string `json:"cause,omitempty"`
}

func (o ShowProcessResqJsonDetails) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProcessResqJsonDetails struct{}"
	}

	return strings.Join([]string{"ShowProcessResqJsonDetails", string(data)}, " ")
}
