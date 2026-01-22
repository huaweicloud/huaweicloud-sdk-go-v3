package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RocketMqExtendProductIosEntity **参数解释**： 磁盘IO信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RocketMqExtendProductIosEntity struct {

	// **参数解释**： 存储IO规格。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IoSpec *string `json:"io_spec,omitempty"`

	// **参数解释**： 可用分区列表。RocketMQ 5.X基础版部署架构为单机时，请选择1个可用区，为集群时可选择1个或2个可用区。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： IO类型。 **约束限制**： 不涉及。 **取值范围**： - evs：[华为云磁盘类型。](tag:hws,hws_hk)[磁盘类型。](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,cmcc,hk_tm,hcs,fcs,dt,hcs_oemout,ax,srg) [- dss：专属云磁盘类型。](tag:hws,hws_hk) **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 不可用分区列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`
}

func (o RocketMqExtendProductIosEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqExtendProductIosEntity struct{}"
	}

	return strings.Join([]string{"RocketMqExtendProductIosEntity", string(data)}, " ")
}
