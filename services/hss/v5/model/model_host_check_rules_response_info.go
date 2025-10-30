package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostCheckRulesResponseInfo 单个主机的单个检查项信息
type HostCheckRulesResponseInfo struct {

	// **参数解释** 检查项id **取值范围** 字符长度0-256位
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释** 检查项（检查规则）名称 **取值范围** 字符长度0-65534位
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// 基线检查中检查项的检查类型 - 访问控制 - 服务配置
	Tag *string `json:"tag,omitempty"`

	// 风险等级，包含如下:   - Low : 低危   - Medium : 中危   - High : 高危
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 检测结果类型 **取值范围** - safe             : 已通过 - unhandled        : 未处理 - ignored          : 已忽略 - fixing           : 修复中 - fix-failed       : 修复失败 - verifying        : 验证中 - add_to_whitelist : 已加白
	ResultType *string `json:"result_type,omitempty"`

	// **参数解释** 支持传递参数修复的检查项可传递参数的范围，只有支持传递参数修复的检查项才返回此数据
	RuleParams *[]CheckRuleFixParamInfo `json:"rule_params,omitempty"`

	// **参数解释** 检查项扫描时间(ms) **取值范围** 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释** 主机类型，是cce则返回cce，否则返回null **取值范围** - cce
	HostType *string `json:"host_type,omitempty"`

	// **参数解释** 差异化展示提示信息 **取值范围** 字符长度0-2048位
	DiffDescription *string `json:"diff_description,omitempty"`

	// **参数解释** 差异化展示提示信息 **取值范围** 字符长度0-512位
	Description *string `json:"description,omitempty"`

	// **参数解释** 是否支持一键修复 **取值范围** - 1 : 支持一键修复 - 0 : 不支持
	EnableFix *int32 `json:"enable_fix,omitempty"`

	// **参数解释** 该检查项的修复 & 验证 按钮是否可单击 **取值范围** - true  : 按钮可单击 - false : 按钮不可单击
	EnableClick *bool `json:"enable_click,omitempty"`

	// **参数解释** 已忽略检查项是否可点击 **取值范围** - true  : 按钮可单击 - false : 按钮不可单击
	CancelIgnoreEnableClick *bool `json:"cancel_ignore_enable_click,omitempty"`

	// **参数解释** 修复失败原因 **取值范围** 不涉及
	FixFailedReason *string `json:"fix_failed_reason,omitempty"`
}

func (o HostCheckRulesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostCheckRulesResponseInfo struct{}"
	}

	return strings.Join([]string{"HostCheckRulesResponseInfo", string(data)}, " ")
}
