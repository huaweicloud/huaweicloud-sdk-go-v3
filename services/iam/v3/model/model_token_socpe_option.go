package model

import (
	"encoding/json"

	"strings"
)

//
type TokenSocpeOption struct {
	Domain *ScopeDomainOption `json:"domain,omitempty"`

	Project *ScopeProjectOption `json:"project,omitempty"`
}

func (o TokenSocpeOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenSocpeOption struct{}"
	}

	return strings.Join([]string{"TokenSocpeOption", string(data)}, " ")
}
