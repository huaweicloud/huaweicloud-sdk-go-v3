package model

import (
	"encoding/json"

	"strings"
)

// scope信息
type GetIdTokenIdScopeBody struct {
	Domain  *GetIdTokenScopeDomainOrProjectBody `json:"domain,omitempty"`
	Project *GetIdTokenScopeDomainOrProjectBody `json:"project,omitempty"`
}

func (o GetIdTokenIdScopeBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetIdTokenIdScopeBody struct{}"
	}

	return strings.Join([]string{"GetIdTokenIdScopeBody", string(data)}, " ")
}
