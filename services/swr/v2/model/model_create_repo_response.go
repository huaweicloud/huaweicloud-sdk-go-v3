package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRepoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRepoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRepoResponse struct{}"
	}

	return strings.Join([]string{"CreateRepoResponse", string(data)}, " ")
}
