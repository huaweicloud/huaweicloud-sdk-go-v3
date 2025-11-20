package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRulesResponseInfo 检查项信息
type CheckRulesResponseInfo struct {

	// **参数解释** 基线检查项的类型 **取值范围** 字符长度0-256位
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 检查项（检查规则）名称 **取值范围** 字符长度0-2048位
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释** 检查项ID **取值范围** 字符长度0-64位
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释** 风险等级，包含如下: **取值范围** - Low    : 低危 - Medium : 中危 - High   : 高危
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 配置检查（基线）的类型,Linux系统支持的基线一般check_type和check_name相同,例如SSH、CentOS 7。 Windows系统支持的基线一般check_type和check_name不相同，例如check_name为Windows的配置检查（基线），它的check_type包含Windows Server 2019 R2、Windows Server 2016 R2等。 **取值范围** 字符长度0-256位
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 配置检查（基线）的类型名称, 一般为check_type + 系统基线检查|应用基线检查 **取值范围** 字符长度0-256位
	CheckTypeName *string `json:"check_type_name,omitempty"`

	// **参数解释** 标准类型，包含如下3种 **取值范围** - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 - cis_standard : 通用安全标准
	Standard *string `json:"standard,omitempty"`

	// **参数解释** 受影响的服务器的数量，进行了当前基线检测的服务器数量 **取值范围** 取值0-2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释** 此检测项失败，且未忽略且未加白的主机数 **取值范围** 取值0-2147483647
	FailedNum *int32 `json:"failed_num,omitempty"`

	// 最新检测时间(ms)
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释** 检查项统计结果： **取值范围** - pass      : 已通过，表示此检查项涉及的主机，全部检查通过。 - failed    : 未通过，表示此检查项涉及的主机，存在检查不通过，且不通过的主机中，存在未处理(忽略、加白)主机。 - processed : 已处理，表示此检查项涉及的主机，存在未通过，但所有的未通过主机均已处理(忽略、加白)。
	StatisticsScanResult *string `json:"statistics_scan_result,omitempty"`

	// **参数解释** 是否支持一键修复 **取值范围** - 1 : 支持一键修复 - 0 : 不支持
	EnableFix *int32 `json:"enable_fix,omitempty"`

	// **参数解释** 该检查项的 修复 & 验证 按钮是否可单击 **取值范围** - true  : 按钮可单击 - false : 按钮不可单击
	EnableClick *bool `json:"enable_click,omitempty"`

	// **参数解释** 已忽略检查项是否可点击 **取值范围** - true  : 按钮可单击 - false : 按钮不可单击
	CancelIgnoreEnableClick *bool `json:"cancel_ignore_enable_click,omitempty"`

	// **参数解释** 支持传递参数修复的检查项可传递参数的范围，只有支持传递参数修复的检查项才返回此数据
	RuleParams *[]CheckRuleFixParamInfo `json:"rule_params,omitempty"`
}

func (o CheckRulesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRulesResponseInfo struct{}"
	}

	return strings.Join([]string{"CheckRulesResponseInfo", string(data)}, " ")
}
