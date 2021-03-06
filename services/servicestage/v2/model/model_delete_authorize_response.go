package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAuthorizeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAuthorizeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAuthorizeResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuthorizeResponse", string(data)}, " ")
}
