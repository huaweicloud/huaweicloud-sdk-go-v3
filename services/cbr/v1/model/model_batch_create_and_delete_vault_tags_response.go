package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateAndDeleteVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateAndDeleteVaultTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateAndDeleteVaultTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsResponse", string(data)}, " ")
}
