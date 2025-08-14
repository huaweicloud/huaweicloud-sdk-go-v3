package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupPolicyRequestInfo1 备份策略
type UpdateBackupPolicyRequestInfo1 struct {

	// **参数解释**: 策略是否启用 **约束限制**: 不涉及 **取值范围**: - true：策略已启用 - false：策略未启用 **默认取值**: true
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 策略ID，若开启防护时开启备份防护，该字段必选 **约束限制**: 不涉及 **取值范围**: 字符长度1-256 **默认取值**: 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	OperationDefinition *OperationDefinitionRequestInfo `json:"operation_definition,omitempty"`

	Trigger *BackupTriggerRequestInfo1 `json:"trigger,omitempty"`
}

func (o UpdateBackupPolicyRequestInfo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyRequestInfo1 struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyRequestInfo1", string(data)}, " ")
}
