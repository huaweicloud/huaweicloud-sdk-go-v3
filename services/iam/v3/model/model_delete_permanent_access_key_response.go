package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePermanentAccessKeyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePermanentAccessKeyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"DeletePermanentAccessKeyResponse", string(data)}, " ")
}
