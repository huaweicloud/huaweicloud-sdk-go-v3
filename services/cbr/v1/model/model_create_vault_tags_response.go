package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVaultTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVaultTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultTagsResponse", string(data)}, " ")
}
