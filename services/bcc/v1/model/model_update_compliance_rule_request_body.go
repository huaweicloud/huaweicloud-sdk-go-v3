package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComplianceRuleRequestBody UpdateComplianceRuleRequestBody
type UpdateComplianceRuleRequestBody struct {

	// 是否启用本地备份
	LocalBackupEnabled bool `json:"local_backup_enabled"`

	// 本地备份副本保留时长
	LocalBackupRetention *int32 `json:"local_backup_retention,omitempty"`

	LocalBackupFrequency *UpdateComplianceRuleRequestBodyLocalBackupFrequency `json:"local_backup_frequency,omitempty"`

	// 是否启用异地复制
	RemoteBackupEnabled *bool `json:"remote_backup_enabled,omitempty"`

	// 异地复制副本保留时长。
	RemoteBackupRetention *int32 `json:"remote_backup_retention,omitempty"`

	// 本地副本是否启用WORM特性。
	LocalWormEnabled *bool `json:"local_worm_enabled,omitempty"`

	// 异地复制副本是否开启WORM特性
	RemoteWormEnabled *bool `json:"remote_worm_enabled,omitempty"`

	// 是否强制开启跨账号备份
	IsCrossAccountBackupForced *bool `json:"is_cross_account_backup_forced,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateComplianceRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComplianceRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateComplianceRuleRequestBody", string(data)}, " ")
}
