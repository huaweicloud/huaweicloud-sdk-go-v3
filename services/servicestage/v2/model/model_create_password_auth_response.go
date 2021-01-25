package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePasswordAuthResponse struct {
	Authorization  *AuthorizationVo `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePasswordAuthResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePasswordAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePasswordAuthResponse", string(data)}, " ")
}
