package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleBackupPolicyInfoResponse Response Object
type ShowSingleBackupPolicyInfoResponse struct {

	// **参数解释**: 策略是否启用 **约束限制**: 不涉及 **取值范围**: true或者false **默认取值**: 不涉及
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 策略ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 策略名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 备份类型。当前包含如下1种。   - backup ：备份 **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	OperationType *string `json:"operation_type,omitempty"`

	OperationDefinition *OperationDefinitionInfo `json:"operation_definition,omitempty"`

	Trigger        *BackupTriggerInfo `json:"trigger,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSingleBackupPolicyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleBackupPolicyInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSingleBackupPolicyInfoResponse", string(data)}, " ")
}
