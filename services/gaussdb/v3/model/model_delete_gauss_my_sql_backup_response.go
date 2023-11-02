package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlBackupResponse Response Object
type DeleteGaussMySqlBackupResponse struct {

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份名称。
	BackupName     *string `json:"backup_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlBackupResponse", string(data)}, " ")
}
