package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBackupRequest Request Object
type DownloadBackupRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 备份ID。
	BackupId string `json:"backup_id"`
}

func (o DownloadBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBackupRequest struct{}"
	}

	return strings.Join([]string{"DownloadBackupRequest", string(data)}, " ")
}
