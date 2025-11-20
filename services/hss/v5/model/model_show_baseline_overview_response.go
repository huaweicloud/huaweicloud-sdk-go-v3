package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineOverviewResponse Response Object
type ShowBaselineOverviewResponse struct {

	// **参数解释**: 最新检测时间(ms)。 **取值范围**: 取值0-9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 检查服务器数量。 **取值范围**: 取值0-2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 未通过主机数。 **取值范围**: 取值0-2147483647
	FailedHostNum *int32 `json:"failed_host_num,omitempty"`

	// **参数解释**: 检测基线数量。例如共检测了SSH、CentOS 7这2个配置检测（基线）类型，值就是2。 **取值范围**: 取值0-2147483647
	CheckTypeNum *int32 `json:"check_type_num,omitempty"`

	// **参数解释**: 基线检测项（检测规则）数量。例如SSH基线检测了17条规则，CentOS 7基线检测了60条规则，值就是17+60=77。 **取值范围**: 取值0-2147483647
	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	// **参数解释**: 基线检查项通过率。 **取值范围**: 取值0-2147483647
	CheckRulePassRate *int32 `json:"check_rule_pass_rate,omitempty"`

	// **参数解释**: 云安全实践基线检查项通过率。 **取值范围**: 取值0-2147483647
	CnStandardCheckRulePassRate *int32 `json:"cn_standard_check_rule_pass_rate,omitempty"`

	// **参数解释**: 等保合规基线检查项通过率。 **取值范围**: 取值0-2147483647
	HwStandardCheckRulePassRate *int32 `json:"hw_standard_check_rule_pass_rate,omitempty"`

	// **参数解释**: 未通过的检查项数量。 **取值范围**: 取值0-2147483647
	CheckRuleFailedNum *int32 `json:"check_rule_failed_num,omitempty"`

	// **参数解释**: 高危检查项数量。 **取值范围**: 取值0-2147483647
	CheckRuleHighRisk *int32 `json:"check_rule_high_risk,omitempty"`

	// **参数解释**: 中危检查项数量。 **取值范围**: 取值0-2147483647
	CheckRuleMediumRisk *int32 `json:"check_rule_medium_risk,omitempty"`

	// **参数解释**: 低危检查项数量。 **取值范围**: 取值0-2147483647
	CheckRuleLowRisk *int32 `json:"check_rule_low_risk,omitempty"`

	// **参数解释**: 弱口令检测主机总数。 **取值范围**: 取值0-2147483647
	WeakPwdTotalHost *int32 `json:"weak_pwd_total_host,omitempty"`

	// **参数解释**: 弱口令检测结果：有弱口令的主机数。 **取值范围**: 取值0-2147483647
	WeakPwdRisk *int32 `json:"weak_pwd_risk,omitempty"`

	// **参数解释**: 弱口令检测结果：无弱口令的主机数。 **取值范围**: 取值0-2147483647
	WeakPwdNormal *int32 `json:"weak_pwd_normal,omitempty"`

	// **参数解释**: 弱口令检测结果：未开启防护的主机数。 **取值范围**: 取值0-2147483647
	WeakPwdNotProtected *int32 `json:"weak_pwd_not_protected,omitempty"`

	// 服务器风险TOP5列表
	HostRisks *[]HostRiskNumInfoResponseInfo `json:"host_risks,omitempty"`

	// 主机弱口令风险TOP5列表
	WeakPwdRiskHosts *[]HostWeakPwdRiskNumInfoResponseInfo `json:"weak_pwd_risk_hosts,omitempty"`
	HttpStatusCode   int                                   `json:"-"`
}

func (o ShowBaselineOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowBaselineOverviewResponse", string(data)}, " ")
}
