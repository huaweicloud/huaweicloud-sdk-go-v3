package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonRisks 节点风险来源
type AddonRisks struct {

	// **参数解释：** 插件模板名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AddonTemplateName *string `json:"addonTemplateName,omitempty"`

	// **参数解释：** 插件别名。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Alias *string `json:"alias,omitempty"`
}

func (o AddonRisks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonRisks struct{}"
	}

	return strings.Join([]string{"AddonRisks", string(data)}, " ")
}
