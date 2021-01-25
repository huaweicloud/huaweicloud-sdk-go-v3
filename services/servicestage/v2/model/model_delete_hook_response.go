package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHookResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteHookResponse struct{}"
	}

	return strings.Join([]string{"DeleteHookResponse", string(data)}, " ")
}
