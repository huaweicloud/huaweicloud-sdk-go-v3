package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreferredAffinity struct {

	// **参数解释**：调度器会优先将Pod调度到满足该字段指定的亲和性表达式的节点上，但它也可能选择违反一个或多个表达式的节点。最优先选择的节点是权重总和最大的节点，即对于每个满足所有调度要求（资源请求、调度期间必需的亲和性表达式等）的节点，通过遍历该字段的元素并计算总和，如果节点匹配相应的匹配表达式，则将“权重”加到总和中；权重总和最高的节点即为最优先选择的节点。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodeAffinity *[]PreferredSchedulingTerm `json:"node_affinity,omitempty"`
}

func (o PreferredAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreferredAffinity struct{}"
	}

	return strings.Join([]string{"PreferredAffinity", string(data)}, " ")
}
