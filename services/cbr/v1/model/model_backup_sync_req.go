package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupSyncReq struct {

	// 待同步备份副本列表
	Sync []BackupSync `json:"sync"`
}

func (o BackupSyncReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSyncReq struct{}"
	}

	return strings.Join([]string{"BackupSyncReq", string(data)}, " ")
}
