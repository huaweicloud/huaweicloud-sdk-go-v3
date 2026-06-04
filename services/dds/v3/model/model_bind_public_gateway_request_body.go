package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindPublicGatewayRequestBody struct {

	// **参数解释：** 公网NAT网关实例的ID。可以调用“查询公网NAT网关列表”接口获取。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NatGatewayId string `json:"nat_gateway_id"`

	// **参数解释：** 弹性公网IP的ID。可以调用“查询弹性公网IP列表”接口获取。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PublicIpId string `json:"public_ip_id"`

	// **参数解释：** 弹性公网IP对外提供服务的端口号。 **约束限制：** 不涉及。 **取值范围：** 1~65535。 **默认取值：** 不涉及。
	ExternalServicePort int32 `json:"external_service_port"`
}

func (o BindPublicGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"BindPublicGatewayRequestBody", string(data)}, " ")
}
