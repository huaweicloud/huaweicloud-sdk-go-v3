package model

import (
	"encoding/json"

	"strings"
)

type VaultBackupReq struct {
	Checkpoint *VaultBackup `json:"checkpoint"`
}

func (o VaultBackupReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultBackupReq struct{}"
	}

	return strings.Join([]string{"VaultBackupReq", string(data)}, " ")
}
