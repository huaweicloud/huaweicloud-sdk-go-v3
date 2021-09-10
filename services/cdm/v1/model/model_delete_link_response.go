package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteLinkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLinkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLinkResponse struct{}"
	}

	return strings.Join([]string{"DeleteLinkResponse", string(data)}, " ")
}
