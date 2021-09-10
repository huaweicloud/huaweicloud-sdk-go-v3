package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDbUserResponse struct {
	DbUser         *DbUser `json:"db_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDbUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDbUserResponse struct{}"
	}

	return strings.Join([]string{"ShowDbUserResponse", string(data)}, " ")
}
