package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFailureJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFailureJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFailureJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteFailureJobResponse", string(data)}, " ")
}
