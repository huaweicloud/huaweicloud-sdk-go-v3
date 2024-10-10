package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteManualBackupResponse Response Object
type DeleteManualBackupResponse struct {

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份名称。
	BackupName     *string `json:"backup_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupResponse", string(data)}, " ")
}
