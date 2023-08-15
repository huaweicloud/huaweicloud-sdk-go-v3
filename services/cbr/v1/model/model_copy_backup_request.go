package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyBackupRequest Request Object
type CopyBackupRequest struct {

	// 复制的备份ID
	BackupId string `json:"backup_id"`

	Body *BackupReplicateReq `json:"body,omitempty"`
}

func (o CopyBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyBackupRequest struct{}"
	}

	return strings.Join([]string{"CopyBackupRequest", string(data)}, " ")
}
