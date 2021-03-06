package model

import (
	"encoding/json"

	"strings"
)

type VaultCreateReq struct {
	Vault *VaultCreate `json:"vault"`
}

func (o VaultCreateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultCreateReq struct{}"
	}

	return strings.Join([]string{"VaultCreateReq", string(data)}, " ")
}
