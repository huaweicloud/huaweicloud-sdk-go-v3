package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindPublicGatewayRequest Request Object
type BindPublicGatewayRequest struct {

	// **参数解释：** 实例ID，可以调用“查询实例列表和详情-QueryingInstancesandDetails”接口获取。如果未申请实例，可以调用“创建实例-CreatingaDBInstance”接口创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 需要绑定公网网关规则的节点ID。 **约束限制：** 集群实例选择mongos节点，副本集实例选择primary或者secondary节点。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NodeId string `json:"node_id"`

	Body *BindPublicGatewayRequestBody `json:"body,omitempty"`
}

func (o BindPublicGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicGatewayRequest struct{}"
	}

	return strings.Join([]string{"BindPublicGatewayRequest", string(data)}, " ")
}
