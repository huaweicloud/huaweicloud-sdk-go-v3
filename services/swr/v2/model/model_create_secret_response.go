package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSecretResponse struct {
	// 认证信息

	Auths          map[string]AuthInfo `json:"auths,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateSecretResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecretResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretResponse", string(data)}, " ")
}
