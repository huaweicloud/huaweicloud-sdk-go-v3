package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleNodePoolSpec 伸缩节点池请求详细参数
type ScaleNodePoolSpec struct {

	// 节点池期望节点数
	DesiredNodeCount int32 `json:"desiredNodeCount"`

	// **参数解释**： 要扩缩容的节点池中的伸缩组名称[，通过[获取指定的节点池](cce_02_0355.xml)接口获取伸缩组名称](tag:hws) **约束限制**： 如果要伸缩默认伸缩组填\"default\" **取值范围**： 不涉及 **默认取值**： 不涉及
	ScaleGroups []string `json:"scaleGroups"`

	Options *ScaleNodePoolOptions `json:"options,omitempty"`
}

func (o ScaleNodePoolSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleNodePoolSpec struct{}"
	}

	return strings.Join([]string{"ScaleNodePoolSpec", string(data)}, " ")
}
