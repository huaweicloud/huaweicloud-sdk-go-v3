package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 优先使用定义的规则调度，且不会影响已经在节点上运行的Pod。即优先选择调度到满足规则的节点，但也可能会调度到不满足规则的节点。
type WeightPodAffinityTerms struct {
	PodAffinityTerm *WeightPodAffinityTermsPodAffinityTerm `json:"podAffinityTerm,omitempty"`
	// 权重，范围为1-100

	Weight *int32 `json:"weight,omitempty"`
}

func (o WeightPodAffinityTerms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeightPodAffinityTerms struct{}"
	}

	return strings.Join([]string{"WeightPodAffinityTerms", string(data)}, " ")
}
