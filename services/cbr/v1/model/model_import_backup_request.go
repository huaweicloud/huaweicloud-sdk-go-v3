package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportBackupRequest Request Object
type ImportBackupRequest struct {
	Body *BackupSyncReq `json:"body,omitempty"`
}

func (o ImportBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackupRequest struct{}"
	}

	return strings.Join([]string{"ImportBackupRequest", string(data)}, " ")
}
