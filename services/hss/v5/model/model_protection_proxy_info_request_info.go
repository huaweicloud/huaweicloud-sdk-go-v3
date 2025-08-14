package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectionProxyInfoRequestInfo **参数解释**: 创建防护策略。若开启勒索防护，新建防护策略，则protection_policy_id为空，create_protection_policy必选 **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
type ProtectionProxyInfoRequestInfo struct {

	// **参数解释**: 策略ID，新建策略可不填。 **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 策略名称，新建防护策略则必填 **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 防护动作，新建防护策略则必填 **约束限制**: 不涉及 **取值范围**: 包含如下：   - alarm_and_isolation ：告警并自动隔离。   - alarm_only ：仅告警。 **默认取值**: 不涉及
	ProtectionMode *string `json:"protection_mode,omitempty"`

	// **参数解释**: 是否开启诱饵防护，新建防护策略则必填。 **约束限制**: 不涉及 **取值范围**: 包含如下1种：   - opened ：开启。 **默认取值**: 不涉及
	BaitProtectionStatus *string `json:"bait_protection_status,omitempty"`

	// **参数解释**: 防护目录，新建防护策略则必填。 **约束限制**: 多个目录请用英文分号隔开，最多支持填写20个防护目录 **取值范围**: 字符长度0-128位，特殊符号只允许使用._-+，不能以空格开头，防护目录长度不得超过256个字符。 **默认取值**: 不涉及
	ProtectionDirectory *string `json:"protection_directory,omitempty"`

	// **参数解释**: 防护文件类型，例如：docx，txt，avi,新建防护策略则必填 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ProtectionType *string `json:"protection_type,omitempty"`

	// **参数解释**: 排除目录(选填) **约束限制**: 多个目录请用英文分号隔开，最多支持填写20个排除目录 **取值范围**: 字符长度0-128位，特殊符号只允许使用._-+，不能以空格开头，防护目录长度不得超过256个字符。 **默认取值**: 不涉及
	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	// **参数解释**: 是否运行时检测 **约束限制**: 不涉及 **取值范围**: 暂时只有关闭一种状态，为预留字段。   - closed ：关闭。 **默认取值**: 不涉及
	RuntimeDetectionStatus *string `json:"runtime_detection_status,omitempty"`

	// **参数解释**: 支持该策略的操作系统，新建防护策略则必填 **约束限制**: 不涉及 **取值范围**: 包含两种：   - Windows : Windows系统   - Linux : Linux系统 **默认取值**: 不涉及
	OperatingSystem *string `json:"operating_system,omitempty"`

	// 进程白名单
	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`
}

func (o ProtectionProxyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionProxyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"ProtectionProxyInfoRequestInfo", string(data)}, " ")
}
