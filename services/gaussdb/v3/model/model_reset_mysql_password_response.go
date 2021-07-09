package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResetMysqlPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetMysqlPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetMysqlPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetMysqlPasswordResponse", string(data)}, " ")
}
