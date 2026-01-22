package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IosEntity struct {

	// **参数解释**： 可用分区列表。RocketMQ 5.X基础版部署架构为单机时，请选择1个可用区，为集群时可选择1个或2个可用区。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 不可用分区列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	// **参数解释**： 存储类型规格编码。 **约束限制**： 不涉及。 **取值范围**： - dms.physical.storage.high.v2：高IO类型磁盘 - dms.physical.storage.ultra.v2：超高IO类型磁盘 [- dms.physical.storage.general：使用通用型SSD的磁盘类型。](tag:hws,hws_hk,dt,ctc,ax) [- dms.physical.storage.extreme：使用极速型SSD的磁盘类型。](tag:hws,hws_hk,dt,ctc,ax) **默认取值**： 不涉及。
	IoSpec *string `json:"io_spec,omitempty"`

	// **参数解释**： 存储所属服务类型。 **约束限制**： 不涉及。 **取值范围**： evs **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o IosEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IosEntity struct{}"
	}

	return strings.Join([]string{"IosEntity", string(data)}, " ")
}
