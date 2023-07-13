package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlBackupRequest Request Object
type DeleteGaussMySqlBackupRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 备份文件ID。
	BackupId string `json:"backup_id"`
}

func (o DeleteGaussMySqlBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlBackupRequest", string(data)}, " ")
}
