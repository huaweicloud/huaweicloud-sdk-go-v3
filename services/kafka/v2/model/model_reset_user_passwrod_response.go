package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResetUserPasswrodResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetUserPasswrodResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodResponse struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodResponse", string(data)}, " ")
}
