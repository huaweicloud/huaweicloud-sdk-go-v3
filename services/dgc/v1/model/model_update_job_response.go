package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}
