package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportBackupResponse Response Object
type ImportBackupResponse struct {

	// 同步备份副本接口的返回信息
	Sync           *[]BackupSyncRespBody `json:"sync,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ImportBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackupResponse struct{}"
	}

	return strings.Join([]string{"ImportBackupResponse", string(data)}, " ")
}
