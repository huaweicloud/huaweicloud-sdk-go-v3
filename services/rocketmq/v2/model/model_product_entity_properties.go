package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductEntityProperties **参数解释**： 产品特性。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ProductEntityProperties struct {

	// **参数解释**： 最大topic数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxTopic *string `json:"max_topic,omitempty"`

	// **参数解释**： broker数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerNum *string `json:"broker_num,omitempty"`

	// **参数解释**： 整个实例的计费核数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Core *string `json:"core,omitempty"`

	// **参数解释**： 实例消费者的最大数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxConsumer *string `json:"max_consumer,omitempty"`

	// **参数解释**： 流量单元，rcu * max_tpc_per_rcu = 规格最大TPS。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Rcu *string `json:"rcu,omitempty"`

	// **参数解释**： 最大存储空间，单位为GB。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxStorage *string `json:"max_storage,omitempty"`

	// **参数解释**： 每个节点的最大存储，单位为GB。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxStoragePerNode *string `json:"max_storage_per_node,omitempty"`

	// **参数解释**： product_id的别名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProductAlias *string `json:"product_alias,omitempty"`

	// **参数解释**： 单个rcu最大TPS。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxTpsPerRcu *string `json:"max_tps_per_rcu,omitempty"`

	// **参数解释**： 消息引擎版本。  **约束限制**： 不涉及。  **取值范围**： [- 4.8.0](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- 5.x](tag:hws,hws_eu,hws_hk,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) **默认取值**： 不涉及。
	EngineVersions *string `json:"engine_versions,omitempty"`

	// **参数解释**： 最小存储空间，单位为GB。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MinStorage *string `json:"min_storage,omitempty"`

	// **参数解释**： 每个节点的最小存储。单位为GB。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MinStoragePerNode *string `json:"min_storage_per_node,omitempty"`

	// **参数解释**： Broker的最大个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxBroker *string `json:"max_broker,omitempty"`

	// **参数解释**： 每个节点最多能创建的Topic个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxTopicPerBroker *string `json:"max_topic_per_broker,omitempty"`

	// **参数解释**： 每个节点的最大消费者数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxConsumerPerBroker *string `json:"max_consumer_per_broker,omitempty"`

	// **参数解释**： Broker的最小个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MinBroker *string `json:"min_broker,omitempty"`
}

func (o ProductEntityProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductEntityProperties struct{}"
	}

	return strings.Join([]string{"ProductEntityProperties", string(data)}, " ")
}
