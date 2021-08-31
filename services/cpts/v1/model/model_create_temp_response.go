package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTempResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTempResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTempResponse struct{}"
	}

	return strings.Join([]string{"CreateTempResponse", string(data)}, " ")
}
