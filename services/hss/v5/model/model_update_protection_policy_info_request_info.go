package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProtectionPolicyInfoRequestInfo struct {

	// **参数解释**: 需要修改的防护策略ID，您可以通过[查询勒索病毒的防护策略列表](ListProtectionPolicy.xml)接口获取ID。 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释**: 策略名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PolicyName string `json:"policy_name"`

	// **参数解释**: 防护动作 **约束限制**: 不涉及 **取值范围**: 包含两种：   - alarm_and_isolation ：告警并自动隔离。   - alarm_only ：仅告警。 **默认取值**: 不涉及
	ProtectionMode string `json:"protection_mode"`

	// **参数解释**: 是否开启诱饵防护 **约束限制**: 不涉及 **取值范围**:   - opened ：开启。 **默认取值**: 不涉及
	BaitProtectionStatus *string `json:"bait_protection_status,omitempty"`

	// **参数解释**: 是否开启动态诱饵 **约束限制**: 不涉及 **取值范围**: 包含两种：   - opened ：开启。   - closed ：关闭。   **默认取值**: closed
	DeployMode *string `json:"deploy_mode,omitempty"`

	// **参数解释**: 防护目录 **约束限制**: 多个目录请用英文分号隔开，最多支持填写20个防护目录 **取值范围**: 字符长度0-128位，特殊符号只允许使用._-+，不能以空格开头，防护目录长度不得超过256个字符。 **默认取值**: 不涉及
	ProtectionDirectory string `json:"protection_directory"`

	// **参数解释**: 防护文件类型，例如：docx，txt，avi **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ProtectionType string `json:"protection_type"`

	// **参数解释**: 排除目录(选填) **约束限制**: 多个目录请用英文分号隔开，最多支持填写20个排除目录 **取值范围**: 字符长度0-128位，特殊符号只允许使用._-+，不能以空格开头，防护目录长度不得超过256个字符。 **默认取值**: 不涉及
	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	// **参数解释**: 开启了此勒索防护策略的agent的id列表 **约束限制**: 不涉及 **取值范围**: 列表最大1000条 **默认取值**: 不涉及
	AgentIdList *[]string `json:"agent_id_list,omitempty"`

	// **参数解释**: 支持该策略的操作系统 **约束限制**: 不涉及 **取值范围**: 包含两种：   - Windows : Windows系统   - Linux : Linux系统 **默认取值**: 不涉及
	OperatingSystem string `json:"operating_system"`

	// **参数解释**: 是否运行时检测 **约束限制**: 不涉及 **取值范围**: 包含如下2种，暂时只有关闭一种状态，为保留字段。   - opened ：开启。   - closed ：关闭。 **默认取值**: 不涉及
	RuntimeDetectionStatus *string `json:"runtime_detection_status,omitempty"`

	// 进程白名单
	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`

	// **参数解释**: 是否开启AI勒索防护，包含如下2种。 **约束限制**: 当前只有Windows系统涉及 **取值范围**: 包含两种：   - opened ：开启。   - closed ：关闭。 **默认取值**: closed
	AiProtectionStatus *string `json:"ai_protection_status,omitempty"`
}

func (o UpdateProtectionPolicyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionPolicyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateProtectionPolicyInfoRequestInfo", string(data)}, " ")
}
