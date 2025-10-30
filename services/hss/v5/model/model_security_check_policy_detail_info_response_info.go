package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckPolicyDetailInfoResponseInfo **参数解释** 基线的详细检查项信息
type SecurityCheckPolicyDetailInfoResponseInfo struct {

	// **参数解释** 检查项唯一值 **取值范围** 字符长度0-256位
	Key *string `json:"key,omitempty"`

	// **参数解释** 检查项id **取值范围** 字符长度0-256位
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释** 检查项（检查规则）名称 **取值范围** 字符长度0-65534位
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释** 检查项类型是否是数值类型 **取值范围** - 1 : 是 - 0 : 不是
	CheckRuleType *int32 `json:"check_rule_type,omitempty"`

	// **参数解释** 配置检查（基线）的类型，例如SSH、CentOS 7、Windows **取值范围** 字符长度0-256位
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 检查项的风险程度，包含如下: **取值范围** - Low    : 低危 - Medium : 中危 - High   : 高危
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 检查项的等级 **取值范围** 字符长度0-32位
	Level *string `json:"level,omitempty"`

	// **参数解释** 检查项是否被选中 **取值范围** - true  : 被选中 - false : 未被选中
	Checked *bool `json:"checked,omitempty"`

	// **参数解释** 可自定义配置的参数
	RuleParams *[]CheckRuleFixParamInfo `json:"rule_params,omitempty"`
}

func (o SecurityCheckPolicyDetailInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckPolicyDetailInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckPolicyDetailInfoResponseInfo", string(data)}, " ")
}
