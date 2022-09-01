package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateManualBackupResponse struct {

	// 手动备份的异步任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 手动备份ID。
	BackupId       *string `json:"backup_id,omitempty" xml:"backup_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}
