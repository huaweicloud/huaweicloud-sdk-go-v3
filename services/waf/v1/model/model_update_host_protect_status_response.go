package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateHostProtectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHostProtectStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusResponse", string(data)}, " ")
}
