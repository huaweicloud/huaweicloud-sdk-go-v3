package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBackupOffsitePolicyRequestBody struct {
	BackupPolicy *UpdateBackupOffsitePolicyInfo `json:"backup_policy"`
}

func (o UpdateBackupOffsitePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupOffsitePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBackupOffsitePolicyRequestBody", string(data)}, " ")
}
