package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDatastoresResponse struct {
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatastoresResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDatastoresResponse struct{}"
	}

	return strings.Join([]string{"ListDatastoresResponse", string(data)}, " ")
}
