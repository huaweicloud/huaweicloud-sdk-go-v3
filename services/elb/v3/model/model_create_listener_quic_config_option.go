package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateListenerQuicConfigOption struct {

	// **参数解释**：监听器关联的QUIC监听器ID。  **约束限制**：指定的listener id必须已存在，且协议类型为QUIC，不能指定为null，否则与enable_quic_upgrade冲突。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt)
	QuicListenerId string `json:"quic_listener_id"`

	// **参数解释**：QUIC升级的开启状态。 开启HTTPS监听器升级QUIC监听器能力。  **约束限制**：不涉及  **取值范围**： - true:开启QUIC升级。 - false：关闭QUIC升级。  **默认取值**：false  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt)
	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty"`
}

func (o CreateListenerQuicConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerQuicConfigOption struct{}"
	}

	return strings.Join([]string{"CreateListenerQuicConfigOption", string(data)}, " ")
}
