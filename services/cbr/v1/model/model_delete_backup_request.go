package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupRequest Request Object
type DeleteBackupRequest struct {

	// 指定删除的备份ID
	BackupId string `json:"backup_id"`
}

func (o DeleteBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupRequest", string(data)}, " ")
}
