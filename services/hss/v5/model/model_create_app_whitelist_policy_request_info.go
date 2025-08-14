package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppWhitelistPolicyRequestInfo 创建策略
type CreateAppWhitelistPolicyRequestInfo struct {

	// 策略名称
	PolicyName string `json:"policy_name"`

	// **参数解释**： 进程白名单策略类型 **取值范围**: - allow：允许指定/授权进程运行 - block：阻止潜在恶意软件运行
	PolicyType string `json:"policy_type"`

	// **参数解释**: 策略学习天数 **取值范围**: 最小值1，最大值1000
	LearningDays int32 `json:"learning_days"`

	// **参数解释**： 是否指定学习目录 **约束限制**： 不涉及 **取值范围**: - true：是 - false：否 **默认取值**： 不涉及
	SpecifiedDir *bool `json:"specified_dir,omitempty"`

	// **参数解释**： 是否开启阻断 **取值范围**: - true：是 - false：否
	Intercept bool `json:"intercept"`

	// **参数解释**： 是否自动开启检测 **取值范围**: - true：是 - false：否
	AutoDetect bool `json:"auto_detect"`

	// 是否自动确认学习结果
	AutoConfirm bool `json:"auto_confirm"`

	// **参数解释**： 主机ID列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HostIdList []string `json:"host_id_list"`

	// **参数解释**： 监控目录列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DirList *[]string `json:"dir_list,omitempty"`

	// **参数解释**： 监控文件后缀名列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ExtList *[]string `json:"ext_list,omitempty"`
}

func (o CreateAppWhitelistPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppWhitelistPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateAppWhitelistPolicyRequestInfo", string(data)}, " ")
}
