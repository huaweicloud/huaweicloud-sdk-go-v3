package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateUserInformationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateUserInformationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserInformationResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserInformationResponse", string(data)}, " ")
}
