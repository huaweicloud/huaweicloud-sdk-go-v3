package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupComplianceRule struct {

	// id
	Id *string `json:"id,omitempty"`

	// 是否系统内置规则
	IsSystemEmbedded *bool `json:"is_system_embedded,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 备份类型
	BackupType *string `json:"backup_type,omitempty"`

	// 是否启用本地备份
	LocalBackupEnabled *bool `json:"local_backup_enabled,omitempty"`

	// 本地备份副本保留时长
	LocalBackupRetention *int32 `json:"local_backup_retention,omitempty"`

	LocalBackupFrequency *BackupComplianceRuleLocalBackupFrequency `json:"local_backup_frequency,omitempty"`

	// 是否启用异地复制。
	RemoteBackupEnabled *bool `json:"remote_backup_enabled,omitempty"`

	// 异地复制副本保留时长。
	RemoteBackupRetention *int32 `json:"remote_backup_retention,omitempty"`

	// 本地副本是否启用WORM特性。
	LocalWormEnabled *bool `json:"local_worm_enabled,omitempty"`

	// 异地复制副本是否启用WORM特性。
	RemoteWormEnabled *bool `json:"remote_worm_enabled,omitempty"`

	// 是否开启强制跨账号备份
	IsCrossAccountBackupForced *bool `json:"is_cross_account_backup_forced,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则创建时间
	Created *int64 `json:"created,omitempty"`

	// 规则更新时间
	Updated *int64 `json:"updated,omitempty"`

	// 合规规则名称
	Name *string `json:"name,omitempty"`
}

func (o BackupComplianceRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupComplianceRule struct{}"
	}

	return strings.Join([]string{"BackupComplianceRule", string(data)}, " ")
}
