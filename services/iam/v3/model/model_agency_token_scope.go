package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyTokenScope struct {
	Domain *AgencyTokenScopeDomain `json:"domain,omitempty"`

	Project *AgencyTokenScopeProject `json:"project,omitempty"`
}

func (o AgencyTokenScope) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyTokenScope struct{}"
	}

	return strings.Join([]string{"AgencyTokenScope", string(data)}, " ")
}
