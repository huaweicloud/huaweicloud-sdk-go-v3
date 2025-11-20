package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupInfo struct {

	// 备份ID。
	Id *string `json:"id,omitempty"`

	// 备份名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 文件大小。
	FileSize *float64 `json:"file_size,omitempty"`

	// 创建时间。
	Created *string `json:"created,omitempty"`

	// 更新时间。
	Updated *string `json:"updated,omitempty"`

	// 备份类型。
	BackupType *string `json:"backup_type,omitempty"`
}

func (o BackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfo struct{}"
	}

	return strings.Join([]string{"BackupInfo", string(data)}, " ")
}
