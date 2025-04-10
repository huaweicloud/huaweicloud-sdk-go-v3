package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleBackupPolicyInfoResponse Response Object
type ShowSingleBackupPolicyInfoResponse struct {

	// 策略是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 策略ID
	Id *string `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 备份类型。当前包含如下1种。   - backup ：备份
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
