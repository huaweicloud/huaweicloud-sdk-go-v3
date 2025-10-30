package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckPolicyGroupInfoResponseInfo 配置检测策略关键信息
type SecurityCheckPolicyGroupInfoResponseInfo struct {

	// **参数解释** 策略组ID **取值范围** 字符长度1-64位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 策略组名称 **取值范围** 字符长度1-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释** 基线检测项（检测规则）数量 **取值范围** 取值1-10000
	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	// **参数解释** 应用服务器数量 **取值范围** 取值0-1000000
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释** 是否支持删除 **取值范围** - true：支持 - false：不支持
	Deletable *bool `json:"deletable,omitempty"`

	// **参数解释** 是否默认策略组 **取值范围** - true：是 - false：否
	DefaultGroup *bool `json:"default_group,omitempty"`

	// **参数解释** 策略支持的操作系统类型 **取值范围** - Linux：linux系统 - Windows：windows系统
	SupportOs *string `json:"support_os,omitempty"`

	PolicyInfo *SecurityCheckPolicyInfoResponseInfo `json:"policy_info"`

	WeakPwdPolicyInfo *SecurityCheckPolicyInfoResponseInfo `json:"weak_pwd_policy_info,omitempty"`

	// **参数解释** 应用的主机的agentID列表
	AgentIdList *[]string `json:"agent_id_list,omitempty"`

	TaskCondition *SecurityCheckTaskCondition `json:"task_condition,omitempty"`

	// **参数解释** 检测周期
	DetectionPeriod *string `json:"detection_period,omitempty"`
}

func (o SecurityCheckPolicyGroupInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckPolicyGroupInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckPolicyGroupInfoResponseInfo", string(data)}, " ")
}
