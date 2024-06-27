package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetNewBackupPolicyRequestBody 设置备份策略请求体
type SetNewBackupPolicyRequestBody struct {
	BackupPolicy *BackupPolicyInfo `json:"backup_policy"`
}

func (o SetNewBackupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetNewBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetNewBackupPolicyRequestBody", string(data)}, " ")
}
