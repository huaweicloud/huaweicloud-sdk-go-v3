package model

import (
	"encoding/json"

	"strings"
)

//
type AuthScope struct {
	Domain *AuthScopeDomain `json:"domain,omitempty"`

	Project *AuthScopeProject `json:"project,omitempty"`
}

func (o AuthScope) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AuthScope struct{}"
	}

	return strings.Join([]string{"AuthScope", string(data)}, " ")
}
