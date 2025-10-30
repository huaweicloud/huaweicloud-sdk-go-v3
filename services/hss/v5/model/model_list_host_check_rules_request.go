package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostCheckRulesRequest Request Object
type ListHostCheckRulesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 主机ID **约束限制** 不涉及 **取值范围** 字符长度0-64位 **默认取值** 不涉及
	HostId string `json:"host_id"`

	// **参数解释** 策略组ID **约束限制** 不涉及 **取值范围** 字符长度0-128位 **默认取值** 不涉及
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释** 检测结果类型 **约束限制** 不涉及 **取值范围** - safe：已通过 - unhandled： 未处理 - ignored：已忽略 - fixing：修复中 - fix-failed：修复失败 - verifying：验证中 - add_to_whitelist：已加白  **默认取值** 不涉及
	ResultType *string `json:"result_type,omitempty"`

	// **参数解释** 基线分类 **约束限制** 不涉及 **取值范围** - cn_standard：等保合规标准 - hw_standard：云安全实践标准 - cis_standard：通用安全标准  **默认取值** 不涉及
	Standard string `json:"standard"`

	// **参数解释** 配置检查（基线）的名称，如SSH、CentOS 7、Windows **约束限制** 不涉及 **取值范围** 不涉及 **默认取值** 不涉及
	CheckName string `json:"check_name"`

	// **参数解释** 检查项（检查规则）名称，支持模糊匹配 **约束限制** 不涉及 **取值范围** 字符长度0-2048位 **默认取值** 不涉及
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释** 基线检查项的类型 **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 风险等级 **约束限制** 不涉及 **取值范围** - Security：安全 - Low：低危 - Medium：中危 - High：高危 - Critical：危急  **默认取值** 不涉及
	Severity *string `json:"severity,omitempty"`
}

func (o ListHostCheckRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostCheckRulesRequest struct{}"
	}

	return strings.Join([]string{"ListHostCheckRulesRequest", string(data)}, " ")
}
