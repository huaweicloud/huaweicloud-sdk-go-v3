package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnginePropertiesEntity **参数解释**： 当前规格实例的属性。
type ListEnginePropertiesEntity struct {

	// **参数解释**： 每个Broker的最大分区数。 **取值范围**： 不涉及。
	MaxPartitionPerBroker *string `json:"max_partition_per_broker,omitempty"`

	// **参数解释**： Broker的最大个数。 **取值范围**： 不涉及。
	MaxBroker *string `json:"max_broker,omitempty"`

	// **参数解释**： 每个节点的最大存储。单位为GB。 **取值范围**： 不涉及。
	MaxStoragePerNode *string `json:"max_storage_per_node,omitempty"`

	// **参数解释**： 每个Broker的最大消费者数。 **取值范围**： 不涉及。
	MaxConsumerPerBroker *string `json:"max_consumer_per_broker,omitempty"`

	// **参数解释**： Broker的最小个数。 **取值范围**： 不涉及。
	MinBroker *string `json:"min_broker,omitempty"`

	// **参数解释**： 每个Broker的最大带宽。 **取值范围**： 不涉及。
	MaxBandwidthPerBroker *string `json:"max_bandwidth_per_broker,omitempty"`

	// **参数解释**： 每个节点的最小存储。单位为GB。 **取值范围**： 不涉及。
	MinStoragePerNode *string `json:"min_storage_per_node,omitempty"`

	// **参数解释**： 每个Broker的最大TPS。 **取值范围**： 不涉及。
	MaxTpsPerBroker *string `json:"max_tps_per_broker,omitempty"`

	// **参数解释**： product_id的别名。 **取值范围**： 不涉及。
	ProductAlias *string `json:"product_alias,omitempty"`
}

func (o ListEnginePropertiesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginePropertiesEntity struct{}"
	}

	return strings.Join([]string{"ListEnginePropertiesEntity", string(data)}, " ")
}
