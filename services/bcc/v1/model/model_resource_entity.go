package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceEntity 资源实体
type ResourceEntity struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源名称
	ResourceName string `json:"resource_name"`

	// 资源类型
	Provider string `json:"provider"`

	// 资源等级名称
	ResourceLevelName *string `json:"resource_level_name,omitempty"`

	// 重要程度
	CriticalLevel *string `json:"critical_level,omitempty"`

	// 是否配置保护
	ConfigProtection bool `json:"config_protection"`

	// 备份合规性
	BackupCompliance *string `json:"backup_compliance,omitempty"`

	// 本地备份存储库名称
	LocalVaultName *string `json:"local_vault_name,omitempty"`

	// 异地备份存储库名称
	RemoteVaultName *string `json:"remote_vault_name,omitempty"`

	// 资源归属区域
	RegionId string `json:"region_id"`

	// 资源创建时间
	CreateTime string `json:"create_time"`

	// 上次盘点时间
	LastInventoryTime string `json:"last_inventory_time"`
}

func (o ResourceEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceEntity struct{}"
	}

	return strings.Join([]string{"ResourceEntity", string(data)}, " ")
}
