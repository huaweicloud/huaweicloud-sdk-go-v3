package model

import (
	"encoding/json"

	"strings"
)

// id token信息
type GetIdTokenIdTokenBody struct {
	// id_token的值

	Id string `json:"id"`
}

func (o GetIdTokenIdTokenBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetIdTokenIdTokenBody struct{}"
	}

	return strings.Join([]string{"GetIdTokenIdTokenBody", string(data)}, " ")
}
