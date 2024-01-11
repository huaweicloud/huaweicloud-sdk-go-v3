package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBackupReq 导出备份请求体
type ExportBackupReq struct {

	// 备份ID
	BackupId string `json:"backup_id"`

	// 导出路径
	ExportPath string `json:"export_path"`
}

func (o ExportBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBackupReq struct{}"
	}

	return strings.Join([]string{"ExportBackupReq", string(data)}, " ")
}
