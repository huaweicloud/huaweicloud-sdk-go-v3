package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectionInfoRequestInfo struct {

	// **参数解释**: 策略支持的操作系统 **约束限制**: 不涉及 **取值范围**: 包含如下：   - Windows : Windows系统   - Linux : Linux系统 **默认取值**: 不涉及
	OperatingSystem string `json:"operating_system"`

	// **参数解释**: 勒索防护是否开启，若开启勒索病毒防护，则必填protection_policy_id或者create_protection_policy其中一项。 **约束限制**: 不涉及 **取值范围**: 包含如下：   - closed ：关闭。   - opened ：开启。 **默认取值**: 不涉及
	RansomProtectionStatus string `json:"ransom_protection_status"`

	// **参数解释**: 勒索防护策略ID，若开启勒索防护，选择已有策略防护，则该字段必选。 **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	ProtectionPolicyId *string `json:"protection_policy_id,omitempty"`

	CreateProtectionPolicy *ProtectionProxyInfoRequestInfo `json:"create_protection_policy,omitempty"`

	// **参数解释**: 是否服务器备份，若选择开启服务器备份，则backup_cycle必填 **约束限制**: 不涉及 **取值范围**: 包含如下：   - closed ：关闭。   - opened ：开启。 **默认取值**: 不涉及
	BackupProtectionStatus string `json:"backup_protection_status"`

	BackupResources *BackupResources `json:"backup_resources,omitempty"`

	// **参数解释**: 备份策略ID **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	BackupCycle *UpdateBackupPolicyRequestInfo1 `json:"backup_cycle,omitempty"`

	// **参数解释**: 开启防护的Agent id列表 **约束限制**: 不涉及 **取值范围**: 列表条数0-64 **默认取值**: 不涉及
	AgentIdList []string `json:"agent_id_list"`

	// **参数解释**: 开启防护的host id列表 **约束限制**: 不涉及 **取值范围**: 列表条数0-64 **默认取值**: 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o ProtectionInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"ProtectionInfoRequestInfo", string(data)}, " ")
}
