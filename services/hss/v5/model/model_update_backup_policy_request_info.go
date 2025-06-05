package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupPolicyRequestInfo 备份策略
type UpdateBackupPolicyRequestInfo struct {

	// **参数解释**: 策略是否启用，缺省值：true **约束限制**: 不涉及 **取值范围**: true或者false **默认取值**: true
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 策略ID,若开启防护时开启备份防护，该字段必选 **约束限制**: 不涉及 **取值范围**: 字符长度1-256 **默认取值**: 不涉及
	PolicyId string `json:"policy_id"`

	OperationDefinition *OperationDefinitionRequestInfo `json:"operation_definition,omitempty"`

	Trigger *BackupTriggerRequestInfo `json:"trigger,omitempty"`
}

func (o UpdateBackupPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyRequestInfo", string(data)}, " ")
}
