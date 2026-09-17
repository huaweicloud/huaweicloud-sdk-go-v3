package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InPlaceRollingUpdate **参数解释：** 原地升级配置。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type InPlaceRollingUpdate struct {

	// **参数解释：** 每批升级的最大节点数量。升级时节点池之间会依次进行升级。节点池内的节点分批升级，第一批升级1个节点，第二批升级2个节点，后续每批升级节点数以2的幂数增加，直到达到您设置的每批最大升级节点数，并会持续作用在下一个节点池中 **约束限制：** 不涉及 **取值范围：** [1-120] **默认取值：** 不涉及
	UserDefinedStep int32 `json:"userDefinedStep"`

	// **参数解释：** 节点升级批次作用域 **约束限制：** 不涉及 **取值范围：** - Cluster：节点升级批次配置应用到整个集群，整个升级过程不重置升级批次 - NodePool：节点升级批次配置应用到节点池，升级每个节点池都会重置升级批次  **默认取值：** Cluster
	Scope *string `json:"scope,omitempty"`
}

func (o InPlaceRollingUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InPlaceRollingUpdate struct{}"
	}

	return strings.Join([]string{"InPlaceRollingUpdate", string(data)}, " ")
}
