package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaResourceEntity struct {

	// 支持rabbitmqInstance、kafkaInstance、rocketmqInstance、tags四种。   - rabbitmqInstance表示RabbitMQ实例配额。   - kafkaInstance表示Kafka实例配额。   - rocketmqInstance表示RocketMQ实例配额。   - tags表示标签的配额。
	Type *string `json:"type,omitempty"`

	// 租户最大可以创建的实例个数，或者每个实例最大可以创建的标签个数。
	Quota *int32 `json:"quota,omitempty"`

	// 已创建的实例个数。
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaResourceEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourceEntity struct{}"
	}

	return strings.Join([]string{"QuotaResourceEntity", string(data)}, " ")
}
