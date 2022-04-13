package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBackupResponse struct {
	Backup         *BackupDetail `json:"backup,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupResponse", string(data)}, " ")
}
