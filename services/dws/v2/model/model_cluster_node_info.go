package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterNodeInfo **参数解释**： 集群节点信息。 **取值范围**： 不涉及。
type ClusterNodeInfo struct {

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 节点名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 节点状态。 **取值范围**： - 100：创建中。 - 199：空闲。 - 200：正常。 - 300：不可用。 - 303：创建失败。 - 304：删除中。 - 400：已删除。 - 800：冻结。 - 900：关机。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 节点子状态。 **取值范围**： - 110：ECS已就绪。 - 120：软件已部署。 - 130：节点已正常。 - 140：节点已监控。
	SubStatus *string `json:"sub_status,omitempty"`

	// **参数解释**： 节点规格。 **取值范围**： 不涉及。
	Spec *string `json:"spec,omitempty"`

	// **参数解释**： 实例创建类型。 **取值范围**： - INST：已使用 - NODE：空闲节点
	InstCreateType *string `json:"inst_create_type,omitempty"`

	// **参数解释**： 节点别名。 **取值范围**： 不涉及。
	AliasName *string `json:"alias_name,omitempty"`

	// **参数解释**： 可用区编码。 **取值范围**： 不涉及。
	AzCode *string `json:"az_code,omitempty"`
}

func (o ClusterNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeInfo struct{}"
	}

	return strings.Join([]string{"ClusterNodeInfo", string(data)}, " ")
}
