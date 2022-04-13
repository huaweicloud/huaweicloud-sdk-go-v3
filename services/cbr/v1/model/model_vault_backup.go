package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultBackup struct {
	Parameters *CheckpointParam `json:"parameters,omitempty"`
	// 存储库ID

	VaultId string `json:"vault_id"`
}

func (o VaultBackup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBackup struct{}"
	}

	return strings.Join([]string{"VaultBackup", string(data)}, " ")
}
