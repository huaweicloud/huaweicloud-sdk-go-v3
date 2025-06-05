package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchStartProtectionRequestInfo struct {

	// 操作系统，包含如下：   - Windows : Windows系统   - Linux : Linux系统
	OperatingSystem string `json:"operating_system"`

	// 勒索防护是否开启，包含如下：   - closed ：关闭。   - opened ：开启。   若选择开启，protection_policy_id必填一项
	RansomProtectionStatus string `json:"ransom_protection_status"`

	// 防护策略ID,若ransom_protection_status为opened,则该字段必选
	ProtectionPolicyId *string `json:"protection_policy_id,omitempty"`

	// 是否服务器备份，包含如下：   - closed ：关闭。   - opened ：开启。   若选择开启服务器备份，则vault_id必填
	BackupProtectionStatus string `json:"backup_protection_status"`

	// 需要绑定的存储库ID，若backup_protection_status为opened，则该字段必填
	VaultId *string `json:"vault_id,omitempty"`

	// 开启防护的host id列表
	HostIdList []string `json:"host_id_list"`
}

func (o BatchStartProtectionRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartProtectionRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchStartProtectionRequestInfo", string(data)}, " ")
}
