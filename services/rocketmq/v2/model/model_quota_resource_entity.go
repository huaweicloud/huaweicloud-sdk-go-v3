package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaResourceEntity struct {

	// **参数解释**： 支持[rabbitmqInstance、](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,cmcc,hk_tm,fcs,ax,srg)kafkaInstance、rocketmqInstance、tags。 [- rabbitmqInstance表示RabbitMQ实例配额。](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,cmcc,hk_tm,fcs,ax,srg) - kafkaInstance表示Kafka实例配额。 - rocketmqInstance表示RocketMQ实例配额。 - tags表示标签的配额。 - kafkaInstancePublic表示kafka公网配额，已弃用。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 租户最大可以创建的实例个数，或者每个实例最大可以创建的标签个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**： 已创建的实例个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaResourceEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourceEntity struct{}"
	}

	return strings.Join([]string{"QuotaResourceEntity", string(data)}, " ")
}
