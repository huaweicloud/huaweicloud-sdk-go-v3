package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Affinity **参数解释：** 节点亲和类型。 **约束限制：** AFFINITY/ANTI_AFFINITY
type Affinity struct {

	// **参数解释：** 节点亲和类型。 **约束限制：** 不涉及。 **取值范围：** - AFFINITY：亲和。 - ANTI_AFFINITY：反亲和。 **默认取值：** 不涉及。
	AffinityType string `json:"affinity_type"`

	// **参数解释：** 是否设置强亲和。 **约束限制：** 不涉及。 **取值范围：** - true：设置强亲和。 - false：不设置强亲和。 **默认取值：** 不涉及。
	Required bool `json:"required"`

	// **参数解释：** 选择节点方式。 **约束限制：** 不涉及。 **取值范围：** IP **默认取值：** 不涉及。
	SelectionMode string `json:"selection_mode"`

	// **参数解释：** 通过上述方式选择的列表，长度不能超过20。 **约束限制：** 不涉及。
	Targets map[string]string `json:"targets"`
}

func (o Affinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Affinity struct{}"
	}

	return strings.Join([]string{"Affinity", string(data)}, " ")
}
