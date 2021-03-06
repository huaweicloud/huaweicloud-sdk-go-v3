package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneUpdateGroupRequestBody struct {
	Group *KeystoneUpdateGroupOption `json:"group"`
}

func (o KeystoneUpdateGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateGroupRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateGroupRequestBody", string(data)}, " ")
}
