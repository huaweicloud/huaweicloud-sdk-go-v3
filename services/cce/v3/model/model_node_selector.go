package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSelector **参数解释：** 节点标签选择器，匹配Kubernetes中nodeSelector相关约束。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type NodeSelector struct {

	// **参数解释：** 标签键 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key string `json:"key"`

	// **参数解释：** 标签值列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Value *[]string `json:"value,omitempty"`

	// **参数解释：** 标签逻辑运算符 **约束限制：** 不涉及 **取值范围：** - in：值在列表中 - notin：值不在列表中 - exists：标签存在 - !：标签不存在 - gt：大于 - lt：小于  **默认取值：** 不涉及
	Operator string `json:"operator"`
}

func (o NodeSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSelector struct{}"
	}

	return strings.Join([]string{"NodeSelector", string(data)}, " ")
}
