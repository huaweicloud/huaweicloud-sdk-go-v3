package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBackupFileRequest struct {
	// 备份记录ID。

	BackupId string `json:"backup_id"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o DeleteBackupFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupFileRequest", string(data)}, " ")
}
