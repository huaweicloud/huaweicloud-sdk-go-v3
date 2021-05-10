package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteClusterTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClusterTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteClusterTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterTagResponse", string(data)}, " ")
}
