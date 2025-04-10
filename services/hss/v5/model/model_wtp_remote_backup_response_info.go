package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WtpRemoteBackupResponseInfo struct {

	// 远端备份地址(ip+port)
	BackupAddr *string `json:"backup_addr,omitempty"`

	// 远端备份主机HostId
	BackupHostId *string `json:"backup_host_id,omitempty"`

	// 远端备份服务器名称
	BackupHostName *string `json:"backup_host_name,omitempty"`

	// 是否开启远端备份
	RemoteBackup *bool `json:"remote_backup,omitempty"`
}

func (o WtpRemoteBackupResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpRemoteBackupResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpRemoteBackupResponseInfo", string(data)}, " ")
}
