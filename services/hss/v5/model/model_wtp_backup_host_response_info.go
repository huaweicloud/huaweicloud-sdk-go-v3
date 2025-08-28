package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WtpBackupHostResponseInfo struct {

	// **参数解释**: 备份服务器IP **取值范围**: 字符长度0-256位
	BackupHostIp *string `json:"backup_host_ip,omitempty"`

	// **参数解释**: 备份服务器端口 **取值范围**: 最小值0，最大值65535
	BackupHostPort *int32 `json:"backup_host_port,omitempty"`

	// **参数解释**: 备份路径 **取值范围**: 字符长度0-256位
	BackupDir *string `json:"backup_dir,omitempty"`

	// **参数解释**: 备份状态 **取值范围**: - -1 ：启动失败 - 0 ：未启动 - 1 ：运行中 - 2 ：启动中
	BackupHostStatus *int32 `json:"backup_host_status,omitempty"`

	// **参数解释**: 远端备份服务器的服务器名称 **取值范围**: 字符长度0-256位
	BackupHostName *string `json:"backup_host_name,omitempty"`

	// **参数解释**: 远端备份服务器的服务器ID **取值范围**: 字符长度0-256位
	BackupHostId *string `json:"backup_host_id,omitempty"`
}

func (o WtpBackupHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpBackupHostResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpBackupHostResponseInfo", string(data)}, " ")
}
