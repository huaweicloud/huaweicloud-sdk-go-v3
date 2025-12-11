package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCheckRuleActionRequest Request Object
type ChangeCheckRuleActionRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 主机ID，不赋值时，查租户所有主机 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位。 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 是否校验cce **约束限制**: 不涉及 **取值范围**: - true：是 - false：否  **默认取值**: 不涉及
	CheckCce *bool `json:"check_cce,omitempty"`

	// **参数解释**: 对检查项的操作类型 **约束限制**: 不涉及 **取值范围**: - ignore：忽略 - unignore：取消忽略 - fix：修复 - verify：验证 - add_to_whitelist：加白  **默认取值**: ignore
	Action string `json:"action"`

	Body *CheckRuleIdListRequestInfo `json:"body,omitempty"`
}

func (o ChangeCheckRuleActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCheckRuleActionRequest struct{}"
	}

	return strings.Join([]string{"ChangeCheckRuleActionRequest", string(data)}, " ")
}
