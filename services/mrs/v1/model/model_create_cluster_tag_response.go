package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateClusterTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateClusterTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateClusterTagResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterTagResponse", string(data)}, " ")
}
