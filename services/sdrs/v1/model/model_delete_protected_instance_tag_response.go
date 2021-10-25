package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteProtectedInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProtectedInstanceTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceTagResponse", string(data)}, " ")
}
