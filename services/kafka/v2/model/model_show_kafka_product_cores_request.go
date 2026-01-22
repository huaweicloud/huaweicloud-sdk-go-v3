package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaProductCoresRequest Request Object
type ShowKafkaProductCoresRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。实例ID非必填项，只有填写实例ID响应体才会返回total_extend_storage_space。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 产品ID。 **约束限制**： 不涉及。 **取值范围**： - s6.2u4g.single.small：规格kafka.2u4g.single.small的产品ID。 - c6.2u4g.single：规格kafka.2u4g.single的产品ID。 [- s6.2u4g.cluster.small：规格kafka.2u4g.cluster.small的产品ID。](tag:hws,hws_hk,hws_eu,dt,ax) - c6.2u4g.cluster：规格kafka.2u4g.cluster的产品ID。 - c6.4u8g.cluster：规格kafka.4u8g.cluster的产品ID。 - c6.8u16g.cluster：规格kafka.8u16g.cluster的产品ID。 - c6.12u24g.cluster：规格kafka.12u24g.cluster的产品ID。 - c6.16u32g.cluster：规格kafka.16u32g.cluster的产品ID。  **默认取值**： 不涉及。
	ProductId string `json:"product_id"`

	// **参数解释**： broker数量。 **约束限制**： 不涉及。 **取值范围**： - Kafka实例规格为kafka.2u4g.single.small时，代理数建议取值范围1。 - Kafka实例规格为kafka.2u4g.single时，代理数建议取值范围1。 [- Kafka实例规格为kafka.2u4g.cluster.small时，代理数取值范围3-30。](tag:hws,hws_hk,hws_eu,dt,ax) - Kafka实例规格为kafka.2u4g.cluster时，代理数取值范围3-30。 - Kafka实例规格为kafka.4u8g.cluster时，代理数取值范围3-30。 - Kafka实例规格为kafka.8u16g.cluster时，代理数取值范围3-50。 - Kafka实例规格为kafka.12u24g.cluster时，代理数取值范围3-50。 - Kafka实例规格为kafka.16u32g.cluster时，代理数取值范围3-50。 **默认取值**： 不涉及。
	BrokerNum string `json:"broker_num"`
}

func (o ShowKafkaProductCoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaProductCoresRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaProductCoresRequest", string(data)}, " ")
}
