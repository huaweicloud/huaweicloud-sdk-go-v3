package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAutoEnlargePoliciesRequestBody struct {

	// **参数解释：** 自动扩容开关。 **约束限制：** 不涉及。 **取值范围：** - on:开启磁盘自动扩容策略。 - off: 关闭磁盘自动扩容策略。 **默认取值：** on。
	SwitchOption *string `json:"switch_option,omitempty"`

	// **参数解释：** 磁盘自动扩容策略。 **约束限制：** 最大支持设置10个实例的策略。 **取值范围：** 不涉及。 **默认取值：** 不涉及。 **参数解释：** 不涉及。
	Policies []DiskSetAutoExpansionPolicy `json:"policies"`
}

func (o SetAutoEnlargePoliciesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePoliciesRequestBody struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePoliciesRequestBody", string(data)}, " ")
}
