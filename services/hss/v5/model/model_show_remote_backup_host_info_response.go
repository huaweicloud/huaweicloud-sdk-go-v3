package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemoteBackupHostInfoResponse Response Object
type ShowRemoteBackupHostInfoResponse struct {

	// 远端备份服务器的地址，包含IP和端口
	BackupAddr *string `json:"backup_addr,omitempty"`

	// 远端备份服务器的服务器ID
	BackupHostId *string `json:"backup_host_id,omitempty"`

	// 远端备份服务器的服务器名称
	BackupHostName *string `json:"backup_host_name,omitempty"`

	// **参数解释** 服务器是否开启远端备份 **取值范围** - true : 已开启远端备份。 - false: 未开启远端备份。
	RemoteBackup   *bool `json:"remote_backup,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowRemoteBackupHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemoteBackupHostInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRemoteBackupHostInfoResponse", string(data)}, " ")
}
