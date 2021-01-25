package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteListenerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteListenerResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerResponse", string(data)}, " ")
}
