package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupPolicyInfoResponse Response Object
type ShowBackupPolicyInfoResponse struct {

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

func (o ShowBackupPolicyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyInfoResponse", string(data)}, " ")
}
