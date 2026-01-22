package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqProductCoresRequest Request Object
type ShowRabbitMqProductCoresRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。实例ID非必填项，只有填写实例ID响应体才会返回total_extend_storage_space。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 产品ID。 **约束限制**： 不涉及。 **取值范围**： - c6.2u4g.single：对应规格rabbitmq.2u4g.single。 - c6.4u8g.single：对应规格rabbitmq.4u8g.single。 - c6.8u16g.single：对应规格rabbitmq.8u16g.single。 - c6.16u32g.single：对应规格rabbitmq.16u32g.single。 - c6.24u48g.single：对应规格rabbitmq.24u48g.single。 - c6.2u4g.cluster：对应规格rabbitmq.2u4g.cluster。 - c6.4u8g.cluster：对应规格rabbitmq.4u8g.cluster。 - c6.8u16g.cluster：对应规格rabbitmq.8u16g.cluster。 - c6.12u24g.cluster：对应规格rabbitmq.12u24g.cluster。 - c6.16u32g.cluster：对应规格rabbitmq.16u32g.cluster。 - c6.24u48g.cluster：对应规格rabbitmq.24u48g.cluster。 - c6.32u64g.cluster：对应规格rabbitmq.32u64g.cluster。 **默认取值**： 不涉及。
	ProductId string `json:"product_id"`

	// **参数解释**： broker数量。 **约束限制**： 不涉及。 **取值范围**： - 1 - 3 - 5 - 7 **默认取值**： 不涉及。
	BrokerNum string `json:"broker_num"`
}

func (o ShowRabbitMqProductCoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProductCoresRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProductCoresRequest", string(data)}, " ")
}
