package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeEngineInstanceReq 变更类型  取值范围： - storage：存储空间扩容，代理数量不变。 - horizontal：代理数量扩容，每个broker的存储空间不变。 - vertical：垂直扩缩容，broker的底层虚机规格变更，代理数量和存储空间不变。
type ResizeEngineInstanceReq struct {

	// 变更类型  取值范围： [- storage：存储空间扩容，代理数量不变。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt) - horizontal：代理数量扩容，每个broker的存储空间不变。 [- vertical：垂直扩容，broker的底层虚机规格变更，代理数量和存储空间不变。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt)
	OperType string `json:"oper_type"`

	// 当oper_type类型是[storage或](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt)horizontal时，该参数有效且必填，实例存储空间 = 代理数量 * 每个broker的存储空间。  [- 当oper_type类型是storage时，代理数量不变，每个broker存储空间最少扩容100GB。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,fcs,dt)  - 当oper_type类型是horizontal时，每个broker的存储空间不变。
	NewStorageSpace *int32 `json:"new_storage_space,omitempty"`

	// 当oper_type类型是vertical时，该参数才有效且必填。
	NewProductId *string `json:"new_product_id,omitempty"`

	// 代理数量  当oper_type参数为horizontal时，该参数必填。
	NewBrokerNum *int32 `json:"new_broker_num,omitempty"`

	// 实例绑定的弹性IP地址的ID。 以英文逗号隔开多个弹性IP地址的ID。 当oper_type参数为horizontal且开启了公网访问时，此参数必填。
	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o ResizeEngineInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceReq", string(data)}, " ")
}
