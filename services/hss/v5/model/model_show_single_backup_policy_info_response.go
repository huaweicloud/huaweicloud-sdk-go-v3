package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleBackupPolicyInfoResponse Response Object
type ShowSingleBackupPolicyInfoResponse struct {

	// **参数解释**: 策略是否启用 **取值范围**: - true：策略已启用 - false：策略未启用
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 策略ID **取值范围**: 字符长度1-128
	Id *string `json:"id,omitempty"`

	// **参数解释**: 策略名称 **取值范围**: 字符长度1-128
	Name *string `json:"name,omitempty"`

	// **参数解释**: 备份类型 **取值范围**: - backup ：备份
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
