package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobResponse", string(data)}, " ")
}
