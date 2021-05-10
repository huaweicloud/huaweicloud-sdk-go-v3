package model

import (
	"encoding/json"

	"strings"
)

type Keypairs struct {
	Keypair *Keypair `json:"keypair"`
}

func (o Keypairs) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Keypairs struct{}"
	}

	return strings.Join([]string{"Keypairs", string(data)}, " ")
}
