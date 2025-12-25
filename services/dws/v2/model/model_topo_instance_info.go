package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopoInstanceInfo **参数解释**： 集群实例信息。 **取值范围**： 不涉及。
type TopoInstanceInfo struct {

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 实例管理IP。 **取值范围**： 不涉及。
	ManageIp *string `json:"manage_ip,omitempty"`

	// **参数解释**： 业务IP。 **取值范围**： 不涉及。
	TrafficIp *string `json:"traffic_ip,omitempty"`

	// **参数解释**： 内部通信IP。 **取值范围**： 不涉及。
	InternalIp *string `json:"internal_ip,omitempty"`

	// **参数解释**： 内部管理IP。 **取值范围**： 不涉及。
	InternalMgntIp *string `json:"internal_mgnt_ip,omitempty"`

	// **参数解释**： 公网IP信息。 **取值范围**： 不涉及。
	Eip *string `json:"eip,omitempty"`

	// **参数解释**： elb地址。 **取值范围**： 不涉及。
	Elb *string `json:"elb,omitempty"`

	// **参数解释**： 节点状态。 **取值范围**： - 100：创建中 - 200：可用 - 300：不可用 - 400：已删除 - 303：创建失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 可用区编码。 **取值范围**： 不涉及。
	AzCode *string `json:"az_code,omitempty"`
}

func (o TopoInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopoInstanceInfo struct{}"
	}

	return strings.Join([]string{"TopoInstanceInfo", string(data)}, " ")
}
