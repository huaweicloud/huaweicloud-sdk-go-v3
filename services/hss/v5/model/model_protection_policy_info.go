package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectionPolicyInfo struct {

	// **参数解释**: 策略ID **取值范围**: 字符长度0-128
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 防护策略名称 **取值范围**: 字符长度1-128
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 防护动作，包含如下2种。   - alarm_and_isolation ：告警并自动隔离。   - alarm_only ：仅告警。 **取值范围**: 字符长度0-128
	ProtectionMode *string `json:"protection_mode,omitempty"`

	// **参数解释**: 是否开启诱饵防护，包含如下1种, 默认为开启防护诱饵防护。   - opened ：开启。   - closed ：关闭。  **取值范围**: 字符长度0-128
	BaitProtectionStatus *string `json:"bait_protection_status,omitempty"`

	// **参数解释**: 是否开启动态诱饵防护，包含如下2种, 默认为关闭动态诱饵防护。   - opened ：开启。   - closed ：关闭。   **取值范围**: 字符长度0-128
	DeployMode *string `json:"deploy_mode,omitempty"`

	// **参数解释**: 防护目录 **取值范围**: 字符长度1-128
	ProtectionDirectory *string `json:"protection_directory,omitempty"`

	// **参数解释**: 防护文件类型，例如：docx，txt，avi **取值范围**: 字符长度1-128
	ProtectionType *string `json:"protection_type,omitempty"`

	// **参数解释**: 排除目录，选填 **取值范围**: 字符长度1-128
	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	// **参数解释**: 是否运行时检测，包含如下2种，暂时只有关闭一种状态，为保留字段。   - opened ：开启。   - closed ：关闭。 **取值范围**: 字符长度0-128
	RuntimeDetectionStatus *string `json:"runtime_detection_status,omitempty"`

	// **参数解释**: 运行时检测目录，现在为保留字段 **取值范围**: 字符长度1-128
	RuntimeDetectionDirectory *string `json:"runtime_detection_directory,omitempty"`

	// **参数解释**: 关联server个数 **取值范围**: 取值范围0-2097152
	CountAssociatedServer *int32 `json:"count_associated_server,omitempty"`

	// **参数解释**: 操作系统类型。 - Linux - Windows **取值范围**: 字符长度1-128
	OperatingSystem *string `json:"operating_system,omitempty"`

	// 进程白名单
	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`

	// **参数解释**: 是否为默认策略，包含如下2种。   - 0 ：非默认策略。   - 1 ：默认策略 **取值范围**: 取值大小0-10
	DefaultPolicy *int32 `json:"default_policy,omitempty"`

	// **参数解释**: 是否开启AI勒索防护，包含如下1种, 默认为开启AI勒索防护。   - opened ：开启。   - closed ：关闭。  **取值范围**: 字符长度1-128
	AiProtectionStatus *string `json:"ai_protection_status,omitempty"`
}

func (o ProtectionPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionPolicyInfo struct{}"
	}

	return strings.Join([]string{"ProtectionPolicyInfo", string(data)}, " ")
}
