package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupRequest Request Object
type UpdateBackupRequest struct {

	// 备份ID
	BackupId string `json:"backup_id"`

	Body *BackupUpdateReq `json:"body,omitempty"`
}

func (o UpdateBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupRequest struct{}"
	}

	return strings.Join([]string{"UpdateBackupRequest", string(data)}, " ")
}
