package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetBackupPolicyRequestBody struct {

	// **参数解释**: 实例ID列表。 **约束限制**: 列表中的实例ID需严格匹配UUID规则，只能由英文字母、数字组成，且长度为36个字符。
	InstanceIds []string `json:"instance_ids"`

	BackupPolicy *BackupPolicyInfo `json:"backup_policy"`
}

func (o BatchSetBackupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSetBackupPolicyRequestBody", string(data)}, " ")
}
