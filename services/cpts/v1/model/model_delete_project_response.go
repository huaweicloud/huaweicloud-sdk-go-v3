package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProjectResponse struct{}"
	}

	return strings.Join([]string{"DeleteProjectResponse", string(data)}, " ")
}
