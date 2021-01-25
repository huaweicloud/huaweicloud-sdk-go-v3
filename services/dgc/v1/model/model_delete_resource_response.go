package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceResponse", string(data)}, " ")
}
