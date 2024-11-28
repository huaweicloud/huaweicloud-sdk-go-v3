package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBackupRequest Request Object
type StopBackupRequest struct {

	// 备份ID。
	BackupId string `json:"backup_id"`

	Body *StopBackupRequestBody `json:"body,omitempty"`
}

func (o StopBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBackupRequest struct{}"
	}

	return strings.Join([]string{"StopBackupRequest", string(data)}, " ")
}
