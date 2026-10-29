package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleItem struct {

	// **参数解释:** 规则ID。 **取值范围:** 32位ID。
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释:** 规则所属语言。 **取值范围:** 当前支持的语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES
	RuleLanguage *string `json:"rule_language,omitempty"`

	// **参数解释:** 规则名称。 **取值范围:** 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释:** 规则问题级别。 **取值范围：** - 0：致命。 - 1：严重。 - 2：一般。 - 3：提示。
	RuleSeverity *string `json:"rule_severity,omitempty"`

	// **参数解释:** 规则标签。 **取值范围:** - huawei - iso5055 - cert - cwe - top10 - top25 - misra - GJB - autosar - PCI DSS
	RuleTages *string `json:"rule_tages,omitempty"`

	// **参数解释:** 规则状态。 **取值范围:** - 0：未启用。 - 1：已启用。
	Checked *string `json:"checked,omitempty"`

	// **参数解释:** 规则配置参数阈值相关信息。
	RuleConfigList *[]RuleConfigParam `json:"rule_config_list,omitempty"`
}

func (o RuleItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleItem struct{}"
	}

	return strings.Join([]string{"RuleItem", string(data)}, " ")
}
