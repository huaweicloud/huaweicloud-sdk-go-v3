package model

import (
	"encoding/json"

	"strings"
)

// 请求体
type GetIdTokenRequestBody struct {
	Auth *GetIdTokenAuthParams `json:"auth"`
}

func (o GetIdTokenRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetIdTokenRequestBody struct{}"
	}

	return strings.Join([]string{"GetIdTokenRequestBody", string(data)}, " ")
}
