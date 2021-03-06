package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteHealthmonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthmonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorResponse", string(data)}, " ")
}
