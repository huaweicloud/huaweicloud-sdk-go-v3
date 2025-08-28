package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBackupHostRequestInfo struct {

	// **参数解释**: 远端备份服务器的服务器ID，如果备份服务器之前不存在，则本次操作为新增，如果之前已存在，则本次操作为修改 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值** : 不涉及
	HostId string `json:"host_id"`

	// **参数解释**: 备份服务器IP **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值** : 不涉及
	BackupHostIp string `json:"backup_host_ip"`

	// **参数解释**: 备份服务器端口 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值65535 **默认取值** : 不涉及
	BackupHostPort int32 `json:"backup_host_port"`

	// **参数解释**: 备份路径 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值** : 不涉及
	BackupDir string `json:"backup_dir"`
}

func (o UpdateBackupHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupHostRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateBackupHostRequestInfo", string(data)}, " ")
}
