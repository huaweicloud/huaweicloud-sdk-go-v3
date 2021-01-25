package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAgencyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgencyResponse", string(data)}, " ")
}
