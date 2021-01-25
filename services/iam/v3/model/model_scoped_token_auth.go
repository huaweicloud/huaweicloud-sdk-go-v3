package model

import (
	"encoding/json"

	"strings"
)

//
type ScopedTokenAuth struct {
	Identity *ScopedTokenIdentity `json:"identity"`
	Scope    *TokenSocpeOption    `json:"scope"`
}

func (o ScopedTokenAuth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ScopedTokenAuth struct{}"
	}

	return strings.Join([]string{"ScopedTokenAuth", string(data)}, " ")
}
