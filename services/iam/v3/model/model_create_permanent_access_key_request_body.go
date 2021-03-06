package model

import (
	"encoding/json"

	"strings"
)

//
type CreatePermanentAccessKeyRequestBody struct {
	Credential *CreateCredentialOption `json:"credential"`
}

func (o CreatePermanentAccessKeyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePermanentAccessKeyRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePermanentAccessKeyRequestBody", string(data)}, " ")
}
