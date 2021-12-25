package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Pod亲和规则
type PodAffinityTerm struct {
	LabelSelector *PodAffinityTermLabelSelector `json:"labelSelector,omitempty"`
	// 命名空间

	Namespaces *[]string `json:"namespaces,omitempty"`
	// 拓扑标签。 先圈定topologyKey指定的范围，然后再选择labelSelector定义的内容，即亲和调度只会在有topologyKey指定的标签节点上调度。

	TopologyKey *string `json:"topologyKey,omitempty"`
}

func (o PodAffinityTerm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodAffinityTerm struct{}"
	}

	return strings.Join([]string{"PodAffinityTerm", string(data)}, " ")
}
