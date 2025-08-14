package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppWhitelistPolicyResponseInfo 策略信息
type AppWhitelistPolicyResponseInfo struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 进程白名单策略类型 **取值范围**: - allow：允许指定/授权进程运行 - block：阻止潜在恶意软件运行
	PolicyType *string `json:"policy_type,omitempty"`

	// **参数解释**： 服务器名称 **约束限制**: 不涉及 **取值范围**: - effecting：学习完成，策略生效 - learned：学习完成，待确认 - learning：学习中 - pause：暂停 - abnormal：学习异常  **默认取值**: 不涉及
	LearningStatus *string `json:"learning_status,omitempty"`

	// **参数解释**: 策略学习天数 **取值范围**: 最小值1，最大值1000
	LearningDays *int32 `json:"learning_days,omitempty"`

	// **参数解释**： 是否指定学习目录 **约束限制**： 不涉及 **取值范围**: - true：是 - false：否 **默认取值**： 不涉及
	SpecifiedDir *bool `json:"specified_dir,omitempty"`

	// 监控目录列表
	DirList *[]string `json:"dir_list,omitempty"`

	// 监控文件后缀名列表
	FileExtensionList *[]string `json:"file_extension_list,omitempty"`

	// **参数解释**： 是否开启阻断 **取值范围**: - true：是 - false：否
	Intercept *bool `json:"intercept,omitempty"`

	// **参数解释**： 是否自动开启检测 **取值范围**: - true：是 - false：否
	AutoDetect *bool `json:"auto_detect,omitempty"`

	// **参数解释**: 学习完成策略未生效主机数 **取值范围**: 最小值0，最大值2147483647
	NotEffectHostNum *int32 `json:"not_effect_host_num,omitempty"`

	// **参数解释**: 学习完成策略已生效主机数 **取值范围**: 最小值0，最大值2147483647
	EffectHostNum *int32 `json:"effect_host_num,omitempty"`

	// **参数解释**: 识别可信进程数 **取值范围**: 最小值0，最大值2147483647
	TrustNum *int32 `json:"trust_num,omitempty"`

	// **参数解释**: 识别可疑进程数 **取值范围**: 最小值0，最大值2147483647
	SuspiciousNum *int32 `json:"suspicious_num,omitempty"`

	// **参数解释**: 识别恶意进程数 **取值范围**: 最小值0，最大值2147483647
	MaliciousNum *int32 `json:"malicious_num,omitempty"`

	// **参数解释**: 识别未知进程数 **取值范围**: 最小值0，最大值2147483647
	UnknownNum *int32 `json:"unknown_num,omitempty"`

	// 学习异常原因列表
	AbnormalInfoList *[]AppWhitelistAbnormalInfo `json:"abnormal_info_list,omitempty"`

	// **参数解释**： 是否自动确认学习结果 **取值范围**: - true：是 - false：否
	AutoConfirm *bool `json:"auto_confirm,omitempty"`

	// **参数解释**： 默认进程白名单策略 **取值范围**: - true：是 - false：否
	DefaultPolicy *bool `json:"default_policy,omitempty"`

	// 主机id集合
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o AppWhitelistPolicyResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppWhitelistPolicyResponseInfo struct{}"
	}

	return strings.Join([]string{"AppWhitelistPolicyResponseInfo", string(data)}, " ")
}
