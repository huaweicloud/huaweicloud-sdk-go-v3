package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckPasswordResponse struct{}"
	}

	return strings.Join([]string{"CheckPasswordResponse", string(data)}, " ")
}
