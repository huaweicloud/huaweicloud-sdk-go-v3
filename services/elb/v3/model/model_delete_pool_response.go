package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePoolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePoolResponse struct{}"
	}

	return strings.Join([]string{"DeletePoolResponse", string(data)}, " ")
}
