package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateGroupRequestBody struct {
	Group *KeystoneCreateGroupOption `json:"group"`
}

func (o KeystoneCreateGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateGroupRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateGroupRequestBody", string(data)}, " ")
}
