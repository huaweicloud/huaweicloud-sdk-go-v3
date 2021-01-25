package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateOAuthResponse struct {
	Authorization  *AuthorizationVo `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateOAuthResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOAuthResponse struct{}"
	}

	return strings.Join([]string{"CreateOAuthResponse", string(data)}, " ")
}
