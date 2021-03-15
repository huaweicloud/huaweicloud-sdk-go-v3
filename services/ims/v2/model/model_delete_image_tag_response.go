package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImageTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteImageTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageTagResponse", string(data)}, " ")
}
