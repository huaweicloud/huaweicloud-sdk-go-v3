package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetNode 重置节点参数。集群内已有节点通过重置进行重新安装并接入集群。
type ResetNode struct {

	// **参数解释**： 节点ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeID string `json:"nodeID"`

	Spec *ReinstallNodeSpec `json:"spec"`
}

func (o ResetNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNode struct{}"
	}

	return strings.Join([]string{"ResetNode", string(data)}, " ")
}
