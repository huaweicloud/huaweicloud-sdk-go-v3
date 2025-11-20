package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRequest Request Object
type ShowBackupRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	// 备份 ID。
	BackupId string `json:"backup_id"`
}

func (o ShowBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupRequest", string(data)}, " ")
}
