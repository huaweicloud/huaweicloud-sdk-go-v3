package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListOsVersionsResponse struct {
	Body           *[]ListOsVersionsResponseBody `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListOsVersionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOsVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListOsVersionsResponse", string(data)}, " ")
}
