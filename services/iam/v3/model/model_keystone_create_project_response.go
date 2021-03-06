package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCreateProjectResponse struct {
	Project        *AuthProjectResult `json:"project,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o KeystoneCreateProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateProjectResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProjectResponse", string(data)}, " ")
}
