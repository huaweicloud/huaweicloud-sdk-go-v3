package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRespClusterBrokers **参数解释**： 节点。
type ShowClusterRespClusterBrokers struct {

	// **参数解释**： 节点IP。 **取值范围**： 不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**： 端口号。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	BrokerId *string `json:"broker_id,omitempty"`

	// **参数解释**： 是否为controller节点。 **取值范围**： - true：是controller节点。 - false：不是controller节点。
	IsController *bool `json:"is_controller,omitempty"`

	// **参数解释**： 服务端版本。 **取值范围**： [- 1.1.0](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,sbc,cmcc,ax) [- 2.3.0](tag:g42,tm,hk_g42,ctc,hk_tm,dt,sbc,cmcc) - 2.7 [- 3.x](tag:hws,hws_hk,dt,sbc,hcs,fcs,ctc,tm,hk_tm,hws_eu,ax)
	Version *string `json:"version,omitempty"`

	// **参数解释**： broker注册时间，为unix时间戳格式。 **取值范围**： 不涉及。
	RegisterTime *int64 `json:"register_time,omitempty"`

	// **参数解释**： Kafka实例节点的连通性是否正常。 **取值范围**： - true：正常。 - false：不正常。
	IsHealth *bool `json:"is_health,omitempty"`
}

func (o ShowClusterRespClusterBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRespClusterBrokers struct{}"
	}

	return strings.Join([]string{"ShowClusterRespClusterBrokers", string(data)}, " ")
}
