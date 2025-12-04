package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowKakfaClusterResponseClusterBrokers struct {

	// **参数解释**： 是否健康。   **取值范围**： - true：健康。 - false：不健康。
	Health *bool `json:"health,omitempty"`

	// **参数解释**： Host地址。   **取值范围**： 不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**： 端口。   **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 节点ID。   **取值范围**： 不涉及。
	BrokerId *string `json:"broker_id,omitempty"`
}

func (o ShowKakfaClusterResponseClusterBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKakfaClusterResponseClusterBrokers struct{}"
	}

	return strings.Join([]string{"ShowKakfaClusterResponseClusterBrokers", string(data)}, " ")
}
