package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RabbitMqExtendProductIosEntity 磁盘IO信息
type RabbitMqExtendProductIosEntity struct {

	// **参数解释**： 存储IO规格。 **取值范围**： - dms.physical.storage.high.v2：高IO云硬盘。 - dms.physical.storage.ultra.v2：超高IO云硬盘。 [- dms.physical.storage.general：通用型SSD云硬盘。](tag:hws,hws_hk,ax,dt) [- dms.physical.storage.extreme：极速型SSD云硬盘。](tag:hws,hws_hk,ax,dt)
	IoSpec *string `json:"io_spec,omitempty"`

	// 有可用资源的可用区列表
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： IO类型。 **取值范围**： evs
	Type *string `json:"type,omitempty"`

	// 资源售罄的可用区列表。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`
}

func (o RabbitMqExtendProductIosEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RabbitMqExtendProductIosEntity struct{}"
	}

	return strings.Join([]string{"RabbitMqExtendProductIosEntity", string(data)}, " ")
}
