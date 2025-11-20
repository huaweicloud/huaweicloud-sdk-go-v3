package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultSecurityCheckPolicyDetailsRequest Request Object
type ShowDefaultSecurityCheckPolicyDetailsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释** 策略支持的操作系统类型 **约束限制** 不涉及 **取值范围** - Linux : Linux操作系统 - Windows : Windows操作系统  **默认取值** Linux
	SupportOs string `json:"support_os"`

	// **参数解释** 配置检查（基线）的类型,例如SSH、CentOS 7、Windows Server 2019 R2、Windows Server 2016 R2、MySQL5-Windows、weakpwd、pwdcomplexity **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 标准类型 **约束限制** 不涉及 **取值范围** - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 - cis_standard: 通用安全标准  **默认取值** 不涉及
	Standard string `json:"standard"`

	// **参数解释** 基线检查项的类型 **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 配置检查（基线）检查项的名称 **约束限制** 不涉及 **取值范围** 字符长度0-2048位 **默认取值** 不涉及
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释** 配置检查（基线）检查项的风险程度 **约束限制** 不涉及 **取值范围** - Low    : 低危 - Medium : 中危 - High   : 高危  **默认取值** 不涉及
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 配置检查（基线）检查项的版本信息 **约束限制** 不涉及 **取值范围** 字符长度0-32位 **默认取值** 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释** 策略组ID **约束限制** 不涉及 **取值范围** 字符长度0-128位 **默认取值** 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 默认是否被选中 **约束限制** 不涉及 **取值范围** false : 不选中 true  : 默认  **默认取值** true
	Checked *bool `json:"checked,omitempty"`
}

func (o ShowDefaultSecurityCheckPolicyDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultSecurityCheckPolicyDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultSecurityCheckPolicyDetailsRequest", string(data)}, " ")
}
