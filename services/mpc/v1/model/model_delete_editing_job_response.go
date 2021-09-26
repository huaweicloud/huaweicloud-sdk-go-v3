package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEditingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEditingJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEditingJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteEditingJobResponse", string(data)}, " ")
}
