package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AffinityNodeAffinity 节点亲和规则
type AffinityNodeAffinity struct {

	// 优先使用定义的规则调度，且不会影响已经在节点上运行的Pod。即优先选择调度到满足规则的节点，但也可能会调度到不满足规则的节点。
	PreferredDuringSchedulingIgnoredDuringExecution *[]PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

	RequiredDuringSchedulingIgnoredDuringExecution *RequiredDuringScheduling `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

func (o AffinityNodeAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AffinityNodeAffinity struct{}"
	}

	return strings.Join([]string{"AffinityNodeAffinity", string(data)}, " ")
}
