package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskSource **参数解释：** 风险项来源。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type RiskSource struct {

	// **参数解释：** 配置风险项。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ConfigurationRisks *[]ConfigurationRisks `json:"configurationRisks,omitempty"`

	// **参数解释：** 废弃API风险。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DeprecatedAPIRisks *[]DeprecatedApiRisks `json:"deprecatedAPIRisks,omitempty"`

	// **参数解释：** 节点风险。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NodeRisks *[]NodeRisks `json:"nodeRisks,omitempty"`

	// **参数解释：** 插件风险。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AddonRisks *[]AddonRisks `json:"addonRisks,omitempty"`
}

func (o RiskSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskSource struct{}"
	}

	return strings.Join([]string{"RiskSource", string(data)}, " ")
}
