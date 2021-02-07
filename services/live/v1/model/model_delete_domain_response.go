package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDomainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDomainResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainResponse", string(data)}, " ")
}
