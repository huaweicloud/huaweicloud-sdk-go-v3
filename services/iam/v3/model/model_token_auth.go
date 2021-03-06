package model

import (
	"encoding/json"

	"strings"
)

//
type TokenAuth struct {
	Identity *TokenAuthIdentity `json:"identity"`
}

func (o TokenAuth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenAuth struct{}"
	}

	return strings.Join([]string{"TokenAuth", string(data)}, " ")
}
