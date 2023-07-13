package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageUsage struct {

	// 备份数量
	BackupCount *int32 `json:"backup_count,omitempty"`

	// 备份容量
	BackupSize *int32 `json:"backup_size,omitempty"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源名称
	ResourceName string `json:"resource_name"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 多AZ备份大小
	BackupSizeMultiaz *int32 `json:"backup_size_multiaz,omitempty"`
}

func (o StorageUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageUsage struct{}"
	}

	return strings.Join([]string{"StorageUsage", string(data)}, " ")
}
