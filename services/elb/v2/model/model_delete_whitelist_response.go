package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWhitelistResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWhitelistResponse struct{}"
	}

	return strings.Join([]string{"DeleteWhitelistResponse", string(data)}, " ")
}
