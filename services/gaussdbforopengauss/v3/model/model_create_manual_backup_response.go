package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualBackupResponse Response Object
type CreateManualBackupResponse struct {
	Backup *BackupInfo `json:"backup,omitempty"`

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}
