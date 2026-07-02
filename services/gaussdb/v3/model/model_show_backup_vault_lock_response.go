package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupVaultLockResponse Response Object
type ShowBackupVaultLockResponse struct {
	BackupVaultlockInfo *BackupVaultLockInfo `json:"backup_vaultlock_info,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowBackupVaultLockResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupVaultLockResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupVaultLockResponse", string(data)}, " ")
}
