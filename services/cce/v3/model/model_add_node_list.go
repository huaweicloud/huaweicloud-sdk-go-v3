package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNodeList 重置节点参数。集群内已有节点，支持通过重置节点方式进行重新安装并接入集群，纳管过程将清理节点上系统盘、数据盘数据，并作为新节点接入Kuberntes集群，请提前备份迁移关键数据。
type AddNodeList struct {

	// **参数解释**： API版本，固定值“v3”。 **约束限制**： 不涉及 **取值范围**： 只能为固定值“v3”。 **默认取值**： 不涉及
	ApiVersion string `json:"apiVersion"`

	// **参数解释**： API类型，，固定值“List”。 **约束限制**： 不涉及 **取值范围**： 只能为固定值“List”。 **默认取值**： 不涉及
	Kind string `json:"kind"`

	// **参数解释**： 纳管节点列表。 **约束限制**： 当前最多支持同时纳管200个节点。
	NodeList []AddNode `json:"nodeList"`
}

func (o AddNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodeList struct{}"
	}

	return strings.Join([]string{"AddNodeList", string(data)}, " ")
}
