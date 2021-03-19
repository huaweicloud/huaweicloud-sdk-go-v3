package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyTokenAuth struct {
	Identity *AgencyTokenIdentity `json:"identity"`

	Scope *AgencyTokenScope `json:"scope"`
}

func (o AgencyTokenAuth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyTokenAuth struct{}"
	}

	return strings.Join([]string{"AgencyTokenAuth", string(data)}, " ")
}
