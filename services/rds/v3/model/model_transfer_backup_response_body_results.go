package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransferBackupResponseBodyResults struct {

	// 备份id
	BackupId *string `json:"backup_id,omitempty"`

	// 任务提交状态
	Status *string `json:"status,omitempty"`

	// 任务id
	JobId *string `json:"job_id,omitempty"`
}

func (o TransferBackupResponseBodyResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferBackupResponseBodyResults struct{}"
	}

	return strings.Join([]string{"TransferBackupResponseBodyResults", string(data)}, " ")
}
