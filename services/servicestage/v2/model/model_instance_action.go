package model

import (
	"encoding/json"

	"strings"
)

type InstanceAction struct {
	Action *InstanceActionType `json:"action"`

	Parameters *InstanceActionParameters `json:"parameters,omitempty"`
}

func (o InstanceAction) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InstanceAction struct{}"
	}

	return strings.Join([]string{"InstanceAction", string(data)}, " ")
}
