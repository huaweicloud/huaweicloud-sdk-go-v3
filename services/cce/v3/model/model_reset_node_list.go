package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetNodeList **参数解释**： 重置节点参数。 > - 满足条件的已有服务器，支持通过纳管节点方式安装并接入集群。 > - 重置过程将清理节点上系统盘、数据盘数据，并作为新节点接入Kuberntes集群，请提前备份迁移关键数据。 > - 其中节点池内节点重置时不支持外部指定配置，将以节点池配置进行校验并重装，以保证同节点池节点一致性。  **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type ResetNodeList struct {

	// **参数解释**： API版本 **约束限制**： 不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释**： API类型 **约束限制**： 不允许修改 **取值范围**： 不涉及 **默认取值**： List
	Kind string `json:"kind"`

	// **参数解释**： API类型 **约束限制**： 不允许修改 **取值范围**： 不涉及 **默认取值**： List
	NodeList []ResetNode `json:"nodeList"`
}

func (o ResetNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodeList struct{}"
	}

	return strings.Join([]string{"ResetNodeList", string(data)}, " ")
}
