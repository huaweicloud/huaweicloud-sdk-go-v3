package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoNodeExpansionPolicyResponse Response Object
type ShowAutoNodeExpansionPolicyResponse struct {

	// **参数解释：** 节点扩容是否开启。 **约束限制：** 不涉及。 **取值范围：**   true为开启。   false为关闭。 **默认取值：** 不涉及。
	SwitchOption *bool `json:"switch_option,omitempty"`

	// **参数解释：** 超负载节点比例。当大等于overload_node_threshold%的节点，节点内存使用率或cpu使用率满足条件时，触发自动扩容节点。 **约束限制：** 不涉及。 **取值范围：** 1-100的正整数。 **默认取值：** 不涉及。
	OverloadNodeThreshold *int32 `json:"overload_node_threshold,omitempty"`

	// **参数解释：** 触发节点自动扩容的CPU使用率。 **约束限制：** 不涉及。 **取值范围：** 1-100的正整数。 **默认取值：** 不涉及。
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`

	// **参数解释：** 触发节点自动扩容的内存使用率。 **约束限制：** 不涉及。 **取值范围：** 1-100的正整数。 **默认取值：** 不涉及。
	MemThreshold *int32 `json:"mem_threshold,omitempty"`

	// **参数解释：** 每次扩容的节点个数。 **约束限制：** 不涉及。 **取值范围：** 大等于1的正整数，最大不超过可扩容的节点上限。 **默认取值：** 不涉及。
	Step *int32 `json:"step,omitempty"`

	// **参数解释：** 自动扩容所能达到的节点上限。 **约束限制：** 不涉及。 **取值范围：** 大等于1的正整数，最大不超过当前实例可扩容的节点上限。 **默认取值：** 不涉及。
	NodeLimit      *int32 `json:"node_limit,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAutoNodeExpansionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoNodeExpansionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoNodeExpansionPolicyResponse", string(data)}, " ")
}
