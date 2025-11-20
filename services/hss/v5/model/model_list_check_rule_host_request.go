package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckRuleHostRequest Request Object
type ListCheckRuleHostRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 具体基线检查项id **约束限制** 不涉及 **取值范围** - 值可以通过这个接口的返回数据获得：/v5/{project_id}/baseline/risk-config/{check_name}/check-rules **默认取值** 不涉及
	CheckRuleId string `json:"check_rule_id"`

	// **参数解释** 配置检查（基线）的名称，如SSH、CentOS 7、Windows，与check_type相比，会多-PID之类的进程信息，通过具体基线维度查询时，传check_name **约束限制** 不涉及 **取值范围** 不涉及 **默认取值** 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释** 配置检查（基线）的类型，如SSH、CentOS 7、Windows，通过检查项维度查询时，传check_type **约束限制** 不涉及 **取值范围** 不涉及 **默认取值** 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 基线分类 **约束限制** 不涉及 **取值范围** - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 - cis_standard：通用安全标准  **默认取值** 不涉及
	Standard string `json:"standard"`

	// **参数解释** 检测结果类型 **约束限制** 不涉及 **取值范围** - safe             : 已通过 - unhandled        : 未处理 - ignored          : 已忽略 - fixing           : 修复中 - fix-failed       : 修复失败 - verifying        : 验证中 - add_to_whitelist : 已加白  **默认取值** 不涉及
	ResultType *string `json:"result_type,omitempty"`

	// **参数解释** 集群ID **约束限制** 不涉及 **取值范围** 字符长度0-64位 **默认取值** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释** 主机名称或ip **约束限制** 不涉及 **取值范围** 不涉及 **默认取值** 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 主机类型，已废弃 **约束限制** 不涉及 **取值范围** - cce **默认取值** 不涉及
	HostType *string `json:"host_type,omitempty"`

	// **参数解释**: 是否只筛选cce主机，已废弃 **约束限制**: 不涉及 **取值范围**: -true：是。 -false：否。 **默认取值**: false
	CheckCce *bool `json:"check_cce,omitempty"`

	// **参数解释** 策略组ID，已废弃 **约束限制** 不涉及 **取值范围** 字符长度0-128位 **默认取值** 不涉及
	PolicyGroupId *string `json:"policy_group_id,omitempty"`
}

func (o ListCheckRuleHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckRuleHostRequest struct{}"
	}

	return strings.Join([]string{"ListCheckRuleHostRequest", string(data)}, " ")
}
