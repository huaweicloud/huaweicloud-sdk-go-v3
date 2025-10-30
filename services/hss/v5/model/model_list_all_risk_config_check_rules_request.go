package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllRiskConfigCheckRulesRequest Request Object
type ListAllRiskConfigCheckRulesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释** 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 标准类型，包含如下: **约束限制** 不涉及 **取值范围** - cn_standard：等保合规标准 - hw_standard：云安全实践标准 - cis_standard：通用安全标准  **默认取值** 不涉及
	Standard *string `json:"standard,omitempty"`

	// **参数解释** 统计结果类型，包含如下： **约束限制** 不涉及 **取值范围** - pass：已通过，表示查看主机全部通过的检查项 - failed：未通过，表示查看主机全部未通过 & 全部未处理的检查项 - processed：已处理，表示查看主机存在未通过 & 未通过主机已全部处理(忽略、加白)的检查项 **默认取值** 不涉及
	StatisticsScanResult *string `json:"statistics_scan_result,omitempty"`

	// **参数解释** 检查项（检查规则）名称，支持模糊匹配 **约束限制** 不涉及 **取值范围** 字符长度0-2048位 **默认取值** 不涉及
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释** 风险等级，包含如下: **约束限制** 不涉及 **取值范围** - Security : 安全 - Low      : 低危 - Medium   : 中危 - High     : 高危 - Critical : 危急 **默认取值** 不涉及
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 集群ID **约束限制** 不涉及 **取值范围** 字符长度0-64位 **默认取值** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释** 基线检查项的类型 **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 策略组ID，不赋值时，查租户所有主机，host_id存在时，此值无效 **约束限制** 不涉及 **取值范围** 字符长度0-128位 **默认取值** 不涉及
	PolicyGroupId *string `json:"policy_group_id,omitempty"`
}

func (o ListAllRiskConfigCheckRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllRiskConfigCheckRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAllRiskConfigCheckRulesRequest", string(data)}, " ")
}
