package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteInstanceUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteInstanceUsersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceUsersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceUsersResponse", string(data)}, " ")
}
