package model

import (
	"encoding/json"

	"strings"
)

//
type MfaTotp struct {
	User *MfaTotpUser `json:"user"`
}

func (o MfaTotp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MfaTotp struct{}"
	}

	return strings.Join([]string{"MfaTotp", string(data)}, " ")
}
