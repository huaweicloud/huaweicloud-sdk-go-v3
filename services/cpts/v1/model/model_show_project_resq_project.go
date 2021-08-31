package model

import (
	"encoding/json"

	"strings"
)

// project
type ShowProjectResqProject struct {
	// create_time

	CreateTime *string `json:"create_time,omitempty"`
	// description

	Description *string `json:"description,omitempty"`
	// group

	Group *string `json:"group,omitempty"`
	// id

	Id *int32 `json:"id,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// source

	Source *int32 `json:"source,omitempty"`
	// update_time

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowProjectResqProject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectResqProject struct{}"
	}

	return strings.Join([]string{"ShowProjectResqProject", string(data)}, " ")
}
