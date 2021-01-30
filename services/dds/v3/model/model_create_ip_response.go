package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateIpResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateIpResponse struct{}"
	}

	return strings.Join([]string{"CreateIpResponse", string(data)}, " ")
}
