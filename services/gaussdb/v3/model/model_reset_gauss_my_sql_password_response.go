package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResetGaussMySqlPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetGaussMySqlPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetGaussMySqlPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetGaussMySqlPasswordResponse", string(data)}, " ")
}
