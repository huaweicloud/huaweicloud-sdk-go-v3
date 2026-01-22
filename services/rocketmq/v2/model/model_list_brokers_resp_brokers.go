package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListBrokersRespBrokers struct {

	// **参数解释**： 全部代理ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ids *[]float32 `json:"ids,omitempty"`

	// **参数解释**： 节点名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerName *string `json:"broker_name,omitempty"`
}

func (o ListBrokersRespBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersRespBrokers struct{}"
	}

	return strings.Join([]string{"ListBrokersRespBrokers", string(data)}, " ")
}
