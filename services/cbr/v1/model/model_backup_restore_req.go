package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 恢复请求body
type BackupRestoreReq struct {
	Restore *BackupRestore `json:"restore"`
}

func (o BackupRestoreReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreReq struct{}"
	}

	return strings.Join([]string{"BackupRestoreReq", string(data)}, " ")
}
