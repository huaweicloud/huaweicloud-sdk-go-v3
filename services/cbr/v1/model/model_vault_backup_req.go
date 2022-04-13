package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultBackupReq struct {
	Checkpoint *VaultBackup `json:"checkpoint"`
}

func (o VaultBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBackupReq struct{}"
	}

	return strings.Join([]string{"VaultBackupReq", string(data)}, " ")
}
