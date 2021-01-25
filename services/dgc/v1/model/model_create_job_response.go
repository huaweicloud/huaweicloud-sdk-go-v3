package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateJobResponse struct{}"
	}

	return strings.Join([]string{"CreateJobResponse", string(data)}, " ")
}
