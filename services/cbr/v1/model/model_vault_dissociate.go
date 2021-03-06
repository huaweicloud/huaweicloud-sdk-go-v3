package model

import (
	"encoding/json"

	"strings"
)

type VaultDissociate struct {
	// 策略ID

	PolicyId string `json:"policy_id"`
}

func (o VaultDissociate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultDissociate struct{}"
	}

	return strings.Join([]string{"VaultDissociate", string(data)}, " ")
}
