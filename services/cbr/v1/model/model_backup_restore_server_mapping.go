package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupRestoreServerMapping struct {

	// 卷备份ID，可以通过控制台查看云服务器备份详情中磁盘级备份的ID；或“查询指定备份”接口，获取备份中children内的磁盘级备份的ID。
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
