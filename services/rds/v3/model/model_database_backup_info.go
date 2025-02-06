package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseBackupInfo struct {

	// 库名
	DatabaseName string `json:"database_name"`

	// 备份文件名
	BackupFileName string `json:"backup_file_name"`

	// 备份文件大小
	BackupFileSize int64 `json:"backup_file_size"`
}

func (o DatabaseBackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseBackupInfo struct{}"
	}

	return strings.Join([]string{"DatabaseBackupInfo", string(data)}, " ")
}
