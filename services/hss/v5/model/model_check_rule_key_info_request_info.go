package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleKeyInfoRequestInfo 检查项key
type CheckRuleKeyInfoRequestInfo struct {

	// **参数解释**: 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**: 检查项ID，值可以通过这个接口的返回数据获得：/v5/{project_id}/baseline/risk-config/{check_name}/check-rules **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释**: 基线标准 **约束限制**: 不涉及 **取值范围**: - cn_standard：等保合规 - hw_standard：云安全实践 - cis_standard：通用安全标准  **默认取值**: 不涉及
	Standard *string `json:"standard,omitempty"`

	// **参数解释**: 用户键入的检查项修复参数数组 **约束限制**: 不涉及 **取值范围**: 元素个数0-10000 **默认取值**: 不涉及
	FixValues *[]CheckRuleFixValuesInfo `json:"fix_values,omitempty"`
}

func (o CheckRuleKeyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleKeyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleKeyInfoRequestInfo", string(data)}, " ")
}
