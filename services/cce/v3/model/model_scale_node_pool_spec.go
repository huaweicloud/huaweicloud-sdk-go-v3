package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleNodePoolSpec 伸缩节点池请求详细参数
type ScaleNodePoolSpec struct {

	// **参数解释**： 节点池的预期总数量。执行扩容操作时，需将当前节点数与扩容数量相加；执行缩容操作时，需从当前节点数中减去缩容数量。 **约束限制**： 必填参数，如果省略则默认值为0，会导致删除节点池伸缩组下的所有节点 **取值范围**： 0或正整数 **默认取值**： 0
	DesiredNodeCount int32 `json:"desiredNodeCount"`

	// **参数解释**： 要扩缩容的节点池中的伸缩组名称，通过[获取指定的节点池](cce_02_0355.xml)接口获取伸缩组名称。扩容时可选择多个伸缩组，系统将按照尽量均分或随机策略在所选伸缩组间分配扩容节点数，缩容时则仅能指定单个伸缩组进行操作。 **约束限制**： 如果要伸缩默认伸缩组填\"default\" **取值范围**： 不涉及 **默认取值**： 不涉及
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
