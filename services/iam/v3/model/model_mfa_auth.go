package model

import (
	"encoding/json"

	"strings"
)

//
type MfaAuth struct {
	Identity *MfaIdentity `json:"identity"`
	Scope    *AuthScope   `json:"scope"`
}

func (o MfaAuth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MfaAuth struct{}"
	}

	return strings.Join([]string{"MfaAuth", string(data)}, " ")
}
