package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 强制使用定义的规则调度，且不会影响已经在节点上运行的Pod。即强制选择调度到满足规则的节点，不会调度到不满足规则的节点。
type RequiredDuringScheduling struct {
	// 节点选择规则

	NodeSelectorTerms *[]MatchExpressions `json:"nodeSelectorTerms,omitempty"`
}

func (o RequiredDuringScheduling) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequiredDuringScheduling struct{}"
	}

	return strings.Join([]string{"RequiredDuringScheduling", string(data)}, " ")
}
