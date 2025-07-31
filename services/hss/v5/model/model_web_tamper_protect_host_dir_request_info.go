package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebTamperProtectHostDirRequestInfo struct {

	// 防护目录
	ProtectDir *string `json:"protect_dir,omitempty"`

	// 排除子目录
	ExcludeChildDir *string `json:"exclude_child_dir,omitempty"`

	// 排除文件路径
	ExcludeFilePath *string `json:"exclude_file_path,omitempty"`

	// 本地备份路径
	LocalBackupDir *string `json:"local_backup_dir,omitempty"`
}

func (o WebTamperProtectHostDirRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectHostDirRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectHostDirRequestInfo", string(data)}, " ")
}
