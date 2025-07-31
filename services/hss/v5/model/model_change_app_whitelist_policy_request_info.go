package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAppWhitelistPolicyRequestInfo 修改策略
type ChangeAppWhitelistPolicyRequestInfo struct {

	// 策略ID
	PolicyId string `json:"policy_id"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// **参数解释**： 进程白名单策略类型进程白名单策略类型 **取值范围**: - allow：允许指定/授权进程运行 - block：阻止潜在恶意软件运行
	PolicyType string `json:"policy_type"`

	// **参数解释**: 策略学习天数 **取值范围**: 最小值1，最大值1000
	LearningDays int32 `json:"learning_days"`

	// **参数解释**： 是否指定学习目录 **约束限制**： 不涉及 **取值范围**: - true：是 - false：否 **默认取值**： 不涉及
	SpecifiedDir *bool `json:"specified_dir,omitempty"`

	// **参数解释**： 是否开启阻断 **取值范围**: - true：是 - false：否
	Intercept bool `json:"intercept"`

	// **参数解释**： 是否自动确认学习结果 **取值范围**: - true：是 - false：否
	AutoConfirm bool `json:"auto_confirm"`

	// **参数解释**： 是否自动开启检测 **取值范围**: - true：是 - false：否
	AutoDetect bool `json:"auto_detect"`

	// 监控目录列表
	DirList *[]string `json:"dir_list,omitempty"`

	// 监控文件后缀名列表
	ExtList *[]string `json:"ext_list,omitempty"`

	// 主机id列表
	HostIdList []string `json:"host_id_list"`
}

func (o ChangeAppWhitelistPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAppWhitelistPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeAppWhitelistPolicyRequestInfo", string(data)}, " ")
}
