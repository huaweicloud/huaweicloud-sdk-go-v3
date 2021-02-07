package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateRepoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRepoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRepoResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepoResponse", string(data)}, " ")
}
