package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyAutoNodeExpansionPolicyRequestBody struct {

	// **参数解释：** 节点扩容是否开启。 **约束限制：** 不涉及。 **取值范围：**   true为开启。   false为关闭。 **默认取值：** 不涉及。
	SwitchOption bool `json:"switch_option"`

	// **参数解释：** 超负载节点比例。如当前实例3个节点，需要当2个节点满足阈值时触发扩容，因2÷3≈67%，可设置为67（向上取整） **约束限制：** 不涉及。 **取值范围：** 1-100的正整数。 **默认取值：** 不涉及。
	OverloadNodeThreshold *int32 `json:"overload_node_threshold,omitempty"`

	// **参数解释：** 触发节点自动扩容的CPU使用率。 **约束限制：** 不涉及。 **取值范围：** 1-100的正整数 **默认取值：** 默认为80。
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`

	// **参数解释：** 触发节点自动扩容的内存使用率。 **约束限制：** 不涉及。 **取值范围：** 1-100的数字 **默认取值：** 默认为80。
	MemThreshold *int32 `json:"mem_threshold,omitempty"`

	// **参数解释：** 每次扩容的节点个数。 **约束限制：** 不涉及。 **取值范围：** 大等于1的正整数，最大不超过可扩容的节点上限。 **默认取值：** 默认为3。
	Step *int32 `json:"step,omitempty"`

	// **参数解释：** 自动扩容所能达到的节点上限。 **约束限制：** 不涉及。 **取值范围：** 大等于1的正整数，最大不超过当前实例可扩容的节点上限。 **默认取值：** 默认为当前实例可扩容的节点上限。
	NodeLimit *int32 `json:"node_limit,omitempty"`
}

func (o ModifyAutoNodeExpansionPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoNodeExpansionPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyAutoNodeExpansionPolicyRequestBody", string(data)}, " ")
}
