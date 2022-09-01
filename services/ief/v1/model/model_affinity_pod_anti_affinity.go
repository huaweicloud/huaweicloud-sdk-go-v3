package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Pod反亲和规则
type AffinityPodAntiAffinity struct {

	// 优先使用定义的规则调度，且不会影响已经在节点上运行的Pod。即优先选择调度到满足规则的节点，但也可能会调度到不满足规则的节点。
	PreferredDuringSchedulingIgnoredDuringExecution *[]WeightPodAffinityTerms `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution"`

	// 强制使用定义的规则调度，且不会影响已经在节点上运行的Pod。即强制选择调度到满足规则的节点，不会调度到不满足规则的节点。
	RequiredDuringSchedulingIgnoredDuringExecution *[]PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

func (o AffinityPodAntiAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AffinityPodAntiAffinity struct{}"
	}

	return strings.Join([]string{"AffinityPodAntiAffinity", string(data)}, " ")
}
