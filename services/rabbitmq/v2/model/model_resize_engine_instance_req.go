package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeEngineInstanceReq struct {

	// 变更类型。  取值范围：  [storage：存储空间扩容，代理数量不变。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)  horizontal：代理数量扩容，每个broker的存储空间不变。  [vertical：垂直扩容，broker的底层虚机规格变更，代理数量和存储空间不变。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)
	OperType string `json:"oper_type"`

	// 扩容后的存储空间。  [当oper_type类型是storage或horizontal时，该参数有效且必填。  实例存储空间 = 代理数量 * 每个broker的存储空间。  当oper_type类型是storage时，代理数量不变，每个broker存储空间最少扩容100GB。  当oper_type类型是horizontal时，每个broker的存储空间不变。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)  [实例存储空间 = 代理数量 * 每个broker的存储空间。 每个broker的存储空间不变。](tag:hcs)
	NewStorageSpace *int32 `json:"new_storage_space,omitempty"`

	// 规格，例如c6.8u16g.cluster，当oper_type类型是vertical时，该参数才有效且必填。
	NewProductId *string `json:"new_product_id,omitempty"`

	// 当oper_type参数为horizontal时，该参数有效。
	NewBrokerNum *int32 `json:"new_broker_num,omitempty"`

	// 老规格，例如dms.instance.rabbitmq.cluster.c3.8u16g，当oper_type类型horizontal时，为dms.instance.rabbitmq.cluster.c3.8u16g.5，最后的数字5为代理数
	NewSpecCode *string `json:"new_spec_code,omitempty"`
}

func (o ResizeEngineInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceReq", string(data)}, " ")
}
