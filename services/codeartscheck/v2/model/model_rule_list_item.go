package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleListItem struct {

	// **参数解释：** 业务类型分类。 **取值范围:** - common：通用。 - common_standard_recommend：通用建议。 - common_standard_rule：通用规则。 - security：安全。 - security_standard_recommend：安全建议。 - security_standard_rule：安全规则。
	RuleSet *string `json:"rule_set,omitempty"`

	// **参数解释:** 规则ID。 **取值范围:** 32位UUID。
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释:** 规则所属语言。 **取值范围:** 当前支持的语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES
	RuleLanguage *string `json:"rule_language,omitempty"`

	// **参数解释:** 规则名称。 **取值范围:** 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释:** 规则问题级别。 **取值范围:** - 0：致命。 - 1：严重。 - 2：一般。 - 3：提示。
	RuleSeverity *string `json:"rule_severity,omitempty"`

	// **参数解释:** 规则标签。 **取值范围:** - huawei - iso5055 - cert - cwe - top10 - top25 - misra - GJB - autosar - PCI DSS
	RuleTages *string `json:"rule_tages,omitempty"`

	// **参数解释:** 正确示例。 **取值范围:** 代码块或者文字示例等。
	RightExample *string `json:"right_example,omitempty"`

	// **参数解释:** 错误示例。 **取值范围:** 代码块或者文字示例等。
	ErrorExample *string `json:"error_example,omitempty"`

	// **参数解释:** 修改建议。 **取值范围:** 代码块或者文字示例等。
	ReviseOpinion *string `json:"revise_opinion,omitempty"`

	// **参数解释:** 规则描述。 **取值范围:** 不涉及。
	RuleDesc *string `json:"rule_desc,omitempty"`
}

func (o RuleListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleListItem struct{}"
	}

	return strings.Join([]string{"RuleListItem", string(data)}, " ")
}
