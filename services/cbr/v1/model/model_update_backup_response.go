package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupResponse Response Object
type UpdateBackupResponse struct {
	Backup         *BackupResp `json:"backup,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupResponse struct{}"
	}

	return strings.Join([]string{"UpdateBackupResponse", string(data)}, " ")
}
