package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskConfigCheckRulesRequest Request Object
type ListRiskConfigCheckRulesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	CheckName string `json:"check_name"`

	// **参数解释**: 标准类型 **约束限制**: 不涉及 **取值范围**: - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 **默认取值**: 不涉及
	Standard string `json:"standard"`

	// **参数解释**: 结果类型 **约束限制**: 不涉及 **取值范围**: - safe : 已通过 - unhandled : 未通过，且未忽略的 - ignored : 未通过，且已忽略的 **默认取值**: unhandled
	ResultType *string `json:"result_type,omitempty"`

	// **参数解释**: 检查项（检查规则）名称，支持模糊匹配 **约束限制**: 不涉及 **取值范围**: 字符长度0-2048位 **默认取值**: 不涉及
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释**: 风险等级 **约束限制**: 不涉及 **取值范围**: - Security : 安全 - Low : 低危 - Medium : 中危 - High : 高危 - Critical : 危急 **默认取值**: 不涉及
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 主机ID，不赋值时，查租户所有主机 **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRiskConfigCheckRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigCheckRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigCheckRulesRequest", string(data)}, " ")
}
