package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskConfigCheckRulesRequest Request Object
type ListRiskConfigCheckRulesRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 配置检查（基线）的名称，例如SSH、CentOS 7、Windows
	CheckName string `json:"check_name"`

	// 标准类型，包含如下: - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准
	Standard string `json:"standard"`

	// 结果类型，包含如下： - safe ： 已通过 - unhandled : 未通过，且未忽略的 - ignored : 未通过，且已忽略的
	ResultType *string `json:"result_type,omitempty"`

	// 检查项（检查规则）名称，支持模糊匹配
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// 风险等级，包含如下:   - Security : 安全   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 危急
	Severity *string `json:"severity,omitempty"`

	// 主机ID，不赋值时，查租户所有主机
	HostId *string `json:"host_id,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRiskConfigCheckRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigCheckRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigCheckRulesRequest", string(data)}, " ")
}
