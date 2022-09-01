package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreBackupRequest struct {

	// 备份id
	BackupId string `json:"backup_id" xml:"backup_id"`

	Body *BackupRestoreReq `json:"body,omitempty" xml:"body"`
}

func (o RestoreBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBackupRequest struct{}"
	}

	return strings.Join([]string{"RestoreBackupRequest", string(data)}, " ")
}
