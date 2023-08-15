package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoSync struct {

	// 同步备份副本数
	SyncBackupNum *int32 `json:"sync_backup_num,omitempty"`

	// 删除的备份副本数
	DeleteBackupNum *int32 `json:"delete_backup_num,omitempty"`

	// 同步失败备副本数
	ErrSyncBackupNum *int32 `json:"err_sync_backup_num,omitempty"`
}

func (o OpExtendInfoSync) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoSync struct{}"
	}

	return strings.Join([]string{"OpExtendInfoSync", string(data)}, " ")
}
