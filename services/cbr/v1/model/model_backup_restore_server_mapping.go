package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupRestoreServerMapping struct {
	// 卷备份ID，可以通过控制台或者“查询指定备份”接口获取。

	BackupId string `json:"backup_id"`
	// 待恢复目标卷ID

	VolumeId string `json:"volume_id"`
}

func (o BackupRestoreServerMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreServerMapping struct{}"
	}

	return strings.Join([]string{"BackupRestoreServerMapping", string(data)}, " ")
}
