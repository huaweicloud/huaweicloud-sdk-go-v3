package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneUpdateProjectRequestBody struct {
	Project *KeystoneUpdateProjectOption `json:"project"`
}

func (o KeystoneUpdateProjectRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProjectRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProjectRequestBody", string(data)}, " ")
}
