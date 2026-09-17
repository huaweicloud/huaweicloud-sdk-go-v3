package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationRisks 配置风险项来源
type ConfigurationRisks struct {

	// **参数解释：** 组件名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Package *string `json:"package,omitempty"`

	// **参数解释：** 涉及文件路径。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SourceFile *string `json:"sourceFile,omitempty"`

	// **参数解释：** 节点信息。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NodeMsg *string `json:"nodeMsg,omitempty"`

	// **参数解释：** 参数值。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Field *string `json:"field,omitempty"`

	// **参数解释：** 修改操作类型。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Operation *string `json:"operation,omitempty"`

	// **参数解释：** 原始值。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	OriginalValue *string `json:"originalValue,omitempty"`

	// **参数解释：** 当前值。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Value *string `json:"value,omitempty"`
}

func (o ConfigurationRisks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRisks struct{}"
	}

	return strings.Join([]string{"ConfigurationRisks", string(data)}, " ")
}
