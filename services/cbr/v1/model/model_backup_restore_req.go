package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 恢复请求body
type BackupRestoreReq struct {
	Restore *BackupRestore `json:"restore" xml:"restore"`
}

func (o BackupRestoreReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreReq struct{}"
	}

	return strings.Join([]string{"BackupRestoreReq", string(data)}, " ")
}
