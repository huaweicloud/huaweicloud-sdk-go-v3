package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeEngineInstanceReq struct {

	// **参数解释**： 变更类型。 **约束限制**： 不涉及。 **取值范围**： [- storage：存储空间扩容，节点数量不变。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt,cmcc,ax,srg) - horizontal：[RocketMQ 5.x为实例规格扩容。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt,srg)[RocketMQ 4.8.0为代理数扩容。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,cmcc,ax) [- vertical：垂直扩容，broker的底层虚机规格变更，代理数量和存储空间不变，仅RocketMQ 4.8.0支持。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,cmcc,ax,srg) **默认取值**： 不涉及。
	OperType string `json:"oper_type"`

	// **参数解释**： 新存储规格 **约束限制**： 当oper_type类型是storage时，该参数有效且必填。  [- 当oper_type类型是storage时，每个broker存储空间最少扩容100GB。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt,cmcc,ax,srg)  - 当oper_type类型是horizontal时，每个broker的存储空间不变。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewStorageSpace *int32 `json:"new_storage_space,omitempty"`

	// **参数解释**： 新产品ID **约束限制**：  [- RocketMQ 5.x：当oper_type类型是horizontal时该参数有效且必填，规格变更仅限于集群实例。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt,srg) [- RocketMQ 4.8.0：当oper_type类型是vertical时该参数有效且必填，规格变更仅限于集群实例。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,cmcc,ax) **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewProductId *string `json:"new_product_id,omitempty"`

	// **参数解释**： 代理数量 **约束限制**： 当oper_type参数为horizontal时，该参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewBrokerNum *int32 `json:"new_broker_num,omitempty"`

	// **参数解释**： 实例绑定的弹性IP地址的ID。以英文逗号隔开多个弹性IP地址的ID。 **约束限制**：  当oper_type参数为horizontal且开启了公网访问时，此参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o ResizeEngineInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceReq", string(data)}, " ")
}
