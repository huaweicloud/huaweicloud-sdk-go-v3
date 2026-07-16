package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequiredAffinityResp 训练作业亲和要求
type RequiredAffinityResp struct {

	// **参数解释**：亲和调度策略。 **取值范围**： - cabinet：强整柜调度 - hyperinstance：超节点亲和调度
	AffinityType *string `json:"affinity_type,omitempty"`

	// **参数解释**：作业整体的网络拓扑约束。 **约束限制**：affinity_type为networkTopology时有效，系统会将作业的所有task调度至不高于job_level层的节点组中。 用户向超节点资源池投递训练作业，如果未设置作业整体的网络拓扑约束，系统会默认赋值为cluster。 **取值范围**： - cluster：资源池级 - hyperinstanceGroup: 超节点级  **默认取值**：默认值cluster。
	JobLevel *string `json:"job_level,omitempty"`

	// **参数解释**：亲和组大小。 **取值范围**：不涉及。
	AffinityGroupSize *int32 `json:"affinity_group_size,omitempty"`

	// **参数解释**：亲和组的网络拓扑约束。 **约束限制**：affinity_type为networkTopology时有效，系统会将affinity_group_size个task组成的亲和组调度至不高于affinity_group_level层的节点组中。 用户向超节点资源池投递训练作业，如果未设置亲和组的网络拓扑约束，系统会默认赋值为hyperinstanceGroup。 **取值范围**： - hyperinstance：超节点级 - slice: 柜级  **默认取值**：默认值hyperinstanceGroup。
	AffinityGroupLevel *string `json:"affinity_group_level,omitempty"`
}

func (o RequiredAffinityResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequiredAffinityResp struct{}"
	}

	return strings.Join([]string{"RequiredAffinityResp", string(data)}, " ")
}
