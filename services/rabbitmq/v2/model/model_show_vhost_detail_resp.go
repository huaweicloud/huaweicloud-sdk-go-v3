package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVhostDetailResp struct {

	// **参数解释**： Vhost名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否开启消息轨迹[（AMQP版本不涉及此字段）](tag:hws,hws_hk,hws_eu,srg)。 **取值范围**： - true：开启。 - false：不开启。
	Tracing *bool `json:"tracing,omitempty"`
}

func (o ShowVhostDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVhostDetailResp struct{}"
	}

	return strings.Join([]string{"ShowVhostDetailResp", string(data)}, " ")
}
