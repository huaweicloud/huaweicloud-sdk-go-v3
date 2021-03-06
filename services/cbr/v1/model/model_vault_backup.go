package model

import (
	"encoding/json"

	"strings"
)

type VaultBackup struct {
	Parameters *CheckpointParam `json:"parameters,omitempty"`
	// 存储库ID

	VaultId string `json:"vault_id"`
}

func (o VaultBackup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultBackup struct{}"
	}

	return strings.Join([]string{"VaultBackup", string(data)}, " ")
}
