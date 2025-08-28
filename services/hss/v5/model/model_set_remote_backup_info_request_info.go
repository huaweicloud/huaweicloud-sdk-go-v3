package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetRemoteBackupInfoRequestInfo struct {

	// **参数解释**: 是否开启远端备份 **约束限制**: 不涉及 **取值范围**: - true ：开启远端备份，需要填写 backup_addr 和 backup_host_id。 - false ：关闭远端备份，无需填写 backup_addr 和 backup_host_id。  **默认取值**: false
	RemoteBackup bool `json:"remote_backup"`

	// **参数解释**: 远端备份地址，需包含IP和端口，格式为IP:端口 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	BackupAddr *string `json:"backup_addr,omitempty"`

	// **参数解释**: 远端备份服务器的服务器ID，填写的远端备份服务器ID必须是已运行中的备份服务器ID **约束限制**: 需要使用 ListBackupHostsInfo 接口查询已运行中的远端备份服务器，在 ListBackupHostsInfo 接口的响应体中，backup_host_status 等于 1 的 backup_host_id 是已运行的远端备份服务器ID **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	BackupHostId *string `json:"backup_host_id,omitempty"`
}

func (o SetRemoteBackupInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRemoteBackupInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"SetRemoteBackupInfoRequestInfo", string(data)}, " ")
}
