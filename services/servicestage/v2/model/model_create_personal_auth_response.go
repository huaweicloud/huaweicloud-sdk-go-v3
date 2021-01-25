package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePersonalAuthResponse struct {
	Authorization  *AuthorizationVo `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePersonalAuthResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePersonalAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePersonalAuthResponse", string(data)}, " ")
}
