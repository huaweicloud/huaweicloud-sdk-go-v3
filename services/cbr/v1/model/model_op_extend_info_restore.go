package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoRestore struct {
	// 备份副本ID

	BackupId *string `json:"backup_id,omitempty"`
	// 备份名称

	BackupName *string `json:"backup_name,omitempty"`
	// 恢复目标资源ID

	TargetResourceId *string `json:"target_resource_id,omitempty"`
	// 恢复目标资源名称

	TargetResourceName *string `json:"target_resource_name,omitempty"`
}

func (o OpExtendInfoRestore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoRestore struct{}"
	}

	return strings.Join([]string{"OpExtendInfoRestore", string(data)}, " ")
}
